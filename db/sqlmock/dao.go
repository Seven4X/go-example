package sqlmock

import "database/sql"

func recordStats(db *sql.DB, userID, productID int64) (err error) {
	tx, err := db.Begin()

	if err != nil {
		return
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	if _, err = tx.Exec("UPDATE products SET views = views + 1"); err != nil {
		return
	}
	if res, err := tx.Exec("INSERT INTO product_viewers (user_id, product_id) VALUES (?, ?)", userID, productID); err != nil {
		println(res)
		return nil
	} else {
		println(res)
	}
	return
}
func recordStatsNoTx(db *sql.DB, userID, productID int64) (err error) {
	_, err = db.Begin()

	if err != nil {
		return
	}

	//defer func() {
	//	switch err {
	//	case nil:
	//		err = tx.Commit()
	//	default:
	//		tx.Rollback()
	//	}
	//}()

	if _, err = db.Exec("UPDATE products SET views = views + 1"); err != nil {
		return
	}
	if res, err := db.Exec("inert into \"topic\" (user_id, product_id) VALUES (?, ?)", userID, productID); err != nil {
		println(res)
		return nil
	} else {
		println(res)
	}
	return
}
