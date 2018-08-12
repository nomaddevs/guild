package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func (db *MariaDB) WriteApplicant(battleid int, battletag, character, email, realName, location, age, gender, computerSpecs,
	previousGuilds, reasonsLeavingGuilds, whyJoinThisGuild, references, finalRemarks string) error {
	conn, err := sql.Open(db.DriverName(), db.ConnectionString())

	if nil != err {
		return err
	}
	defer conn.Close()

	statement := `INSERT INTO applications(battleid, status, battletag, wowcharacter, email, realname, location,
	age, gender, computerspecs, previousguilds, reasonsleavingguilds, whyjointhisguild, 
	wowreferences, finalremarks) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	in, err := conn.Prepare(statement)
	if err != nil {
		return err
	}
	defer in.Close()

	in.Exec(battleid, 2, battletag, character, email, realName, location, age, gender, computerSpecs,
		previousGuilds, reasonsLeavingGuilds, whyJoinThisGuild, references, finalRemarks)

	return nil
}

func (db *MariaDB) GetApplicant(id int) (bool, error) {
	conn, err := sql.Open(db.DriverName(), db.ConnectionString())

	if nil != err {
		return false, err
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM applications WHERE id = ?", id)

	if nil != err {
		return false, err
	}
	defer rows.Close()

	count := 0

	for rows.Next() {
		count++
	}

	err = rows.Err()

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// AcceptApplicant accepts an applicant by setting their status to 1(Accepted).
func (db *MariaDB) AcceptApplicant(id int) error {
	conn, err := sql.Open(db.DriverName(), db.ConnectionString())

	if nil != err {
		return err
	}
	defer conn.Close()

	in, err := conn.Prepare("UPDATE applications SET status = 1 WHERE id = ?")
	if err != nil {
		return err
	}
	defer in.Close()

	in.Exec(id)

	return nil
}

// RejectApplicant rejects an applicant by setting their status to 0(Rejected).
func (db *MariaDB) RejectApplicant(id int) error {
	conn, err := sql.Open(db.DriverName(), db.ConnectionString())

	if nil != err {
		return err
	}
	defer conn.Close()

	in, err := conn.Prepare("UPDATE applications SET status = 0 WHERE id = ?")
	if err != nil {
		return err
	}
	defer in.Close()

	in.Exec(id)

	return nil
}

// PurgeApplicant purges an applicant from the database by BattleID.
func (db *MariaDB) PurgeApplicant(battleid int) error {
	conn, err := sql.Open(db.DriverName(), db.ConnectionString())

	if nil != err {
		return err
	}
	defer conn.Close()

	in, err := conn.Prepare("DELETE FROM applications WHERE battleid = ?")
	if err != nil {
		return err
	}
	defer in.Close()

	in.Exec(battleid)

	return nil
}

func (db *MariaDB) ViewApplicant(id int) ([]int, []string, []string, []string, []string, []string, []string, []string, []string,
	[]string, []string, []string, []string, []string, error) {
	var (
		a int
		b string
		c string
		d string
		e string
		f string
		g string
		h string
		i string
		j string
		k string
		l string
		m string
		n string

		as []int
		bs []string
		cs []string
		ds []string
		es []string
		fs []string
		gs []string
		hs []string
		is []string
		js []string
		ks []string
		ls []string
		ms []string
		ns []string
	)

	conn, err := sql.Open(db.DriverName(), db.ConnectionString())

	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM applications WHERE id = ?", id)

	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&a, &b, &c, &d, &e, &f, &g, &h, &i, &j, &k, &l, &m, &n)

		if err != nil {
			return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
		}

		as = append(as, a)
		bs = append(bs, b)
		cs = append(cs, c)
		ds = append(ds, d)
		es = append(es, e)
		fs = append(fs, f)
		gs = append(gs, g)
		hs = append(hs, h)
		is = append(is, i)
		js = append(js, j)
		ks = append(ks, k)
		ls = append(ls, l)
		ms = append(ms, m)
		ns = append(ns, n)
	}

	err = rows.Err()

	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}

	return as, bs, cs, ds, es, fs, gs, hs, is, js, ks, ls, ms, ns, nil
}

func (db *MariaDB) ViewAllApplicants() ([]int, []string, []string, []string, []string, []string, []string, []string, []string,
	[]string, []string, []string, []string, []string, error) {
	var (
		a int
		b string
		c string
		d string
		e string
		f string
		g string
		h string
		i string
		j string
		k string
		l string
		m string
		n string

		as []int
		bs []string
		cs []string
		ds []string
		es []string
		fs []string
		gs []string
		hs []string
		is []string
		js []string
		ks []string
		ls []string
		ms []string
		ns []string
	)

	conn, err := sql.Open(db.DriverName(), db.ConnectionString())

	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT * FROM applications")

	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&a, &b, &c, &d, &e, &f, &g, &h, &i, &j, &k, &l, &m, &n)

		if err != nil {
			return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
		}

		as = append(as, a)
		bs = append(bs, b)
		cs = append(cs, c)
		ds = append(ds, d)
		es = append(es, e)
		fs = append(fs, f)
		gs = append(gs, g)
		hs = append(hs, h)
		is = append(is, i)
		js = append(js, j)
		ks = append(ks, k)
		ls = append(ls, l)
		ms = append(ms, m)
		ns = append(ns, n)
	}

	err = rows.Err()

	if err != nil {
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, err
	}

	return as, bs, cs, ds, es, fs, gs, hs, is, js, ks, ls, ms, ns, nil
}
