package imysql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/go-sql-driver/mysql"
	"github.com/qjpcpu/ethereum/ethnonce"
	"strings"
	"time"
)

type mysqlManager struct {
	Table   string
	ethConn *ethclient.Client
	db      *sql.DB
}

type MysqlManagerCreator struct {
	mgr *mysqlManager
}

type nonceRecord struct {
	Id       uint64
	Address  string
	Nonce    uint64
	Commit   int
	LastGive int64
}

func (rc *MysqlManagerCreator) SetEthClient(conn *ethclient.Client) ethnonce.ManagerCreator {
	rc.mgr.ethConn = conn
	return rc
}

func (rc *MysqlManagerCreator) Build() *ethnonce.NonceManager {
	return &ethnonce.NonceManager{
		Impl: rc.mgr,
	}
}

func PrepareMysqlManager(connection_str string, tablename string) ethnonce.ManagerCreator {
	db, err := sql.Open("mysql", connection_str)
	if err != nil {
		panic(err)
	}
	stmt, _ := db.Prepare("desc " + "`" + tablename + "`")
	if _, err := stmt.Exec(); err != nil && strings.Contains(err.Error(), "doesn't exist") {
		fmt.Printf("auto create tabel `%s`\n", tablename)
		stmt, err = db.Prepare(fmt.Sprintf(`CREATE TABLE %s (
  id bigint(11) unsigned NOT NULL AUTO_INCREMENT,
  address varchar(42) DEFAULT '' COMMENT 'eth address',
  nonce bigint(20) DEFAULT '0' COMMENT 'nonce',
  commit tinyint(4) DEFAULT '0' COMMENT 'is commited',
  last_give bigint(20) DEFAULT '0' COMMENT 'request give at',
  updated_at datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'modify time',
  PRIMARY KEY (id),
  UNIQUE KEY address (address)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;`, "`"+tablename+"`"))
		stmt.Exec()
	}
	return &MysqlManagerCreator{
		mgr: &mysqlManager{
			db:    db,
			Table: tablename,
		},
	}
}

func (n *mysqlManager) escapedTable() string {
	return "`" + n.Table + "`"
}

func (n *mysqlManager) getRecord(addr common.Address) (nonceRecord, error) {
	row := n.db.QueryRow("SELECT id,address,nonce,commit,last_give FROM "+n.escapedTable()+" WHERE address=?", strings.ToLower(addr.Hex()))
	record := nonceRecord{}
	err := row.Scan(&record.Id, &record.Address, &record.Nonce, &record.Commit, &record.LastGive)
	return record, err
}

func (n *mysqlManager) PeekNonce(addr common.Address) uint64 {
	r, _ := n.getRecord(addr)
	return r.Nonce
}

func (n *mysqlManager) GiveNonce(addr common.Address) (uint64, error) {
	rec, err := n.getRecord(addr)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, ethnonce.ErrNotInitAddress
		}
		return 0, err
	}
	if rec.Commit == 0 {
		stmt, err := n.db.Prepare("UPDATE " + n.escapedTable() + " SET commit=?,last_give=? WHERE address=? AND nonce=?")
		if err != nil {
			return 0, err
		}
		res, err := stmt.Exec(1, time.Now().Unix(), strings.ToLower(addr.Hex()), rec.Nonce)
		if err != nil {
			return 0, err
		}
		if af, _ := res.RowsAffected(); af != 1 {
			return 0, ethnonce.ErrOtherHoldNonce
		}
		return rec.Nonce, nil
	}
	if time.Unix(rec.LastGive, 0).Add(time.Second * 60).Before(time.Now()) {
		stmt, err := n.db.Prepare("UPDATE " + n.escapedTable() + " SET commit=?,last_give=? WHERE address=? AND nonce=? AND last_give=?")
		if err != nil {
			return 0, err
		}
		res, err := stmt.Exec(1, time.Now().Unix(), strings.ToLower(addr.Hex()), rec.Nonce, rec.LastGive)
		if err != nil {
			return 0, err
		}
		if af, _ := res.RowsAffected(); af != 1 {
			return 0, ethnonce.ErrOtherHoldNonce
		}
		return rec.Nonce, nil
	}

	return 0, ethnonce.ErrOtherHoldNonce
}

func (n *mysqlManager) SyncNonce(addr common.Address) (uint64, error) {
	nonce, err := n.ethConn.PendingNonceAt(context.Background(), addr)
	if err != nil {
		return 0, err
	}
	rec, err := n.getRecord(addr)
	if err != nil {
		stmt, err := n.db.Prepare("INSERT INTO " + n.escapedTable() + " (address,nonce,commit,last_give) VALUES(?,?,?,?)")
		if err != nil {
			return 0, err
		}
		stmt.Exec(strings.ToLower(addr.Hex()), nonce, 0, time.Now().Unix())
	} else {
		if rec.Commit == 1 && time.Unix(rec.LastGive, 0).Add(60*time.Second).After(time.Now()) {
			return 0, ethnonce.ErrOtherHoldNonce
		}
		stmt, err := n.db.Prepare("UPDATE " + n.escapedTable() + " SET commit=0,last_give=? WHERE address=? AND nonce=?")
		if err != nil {
			return 0, err
		}
		stmt.Exec(0, strings.ToLower(addr.Hex()), nonce)
	}
	return nonce, err
}

func (n *mysqlManager) CommitNonce(addr common.Address, nonce_number uint64, success bool) error {
	rec, err := n.getRecord(addr)
	if err != nil {
		return ethnonce.ErrNotInitAddress
	}
	if rec.Commit == 0 {
		return nil
	}
	if rec.Nonce != nonce_number {
		return ethnonce.ErrOtherHoldNonce
	}
	var delta int64
	if success {
		delta = 1
	} else {
		delta = 0
	}
	stmt, err := n.db.Prepare("UPDATE " + n.escapedTable() + " SET commit=0,last_give=0,nonce=nonce+? WHERE address=? AND nonce=?")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(delta, strings.ToLower(addr.Hex()), rec.Nonce)
	return err
}

func (n *mysqlManager) Close() error {
	return n.db.Close()
}
