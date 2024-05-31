// DO NOT EDIT THIS FILE!
// It is autogenerated by gentity

package main

import (
	"fmt"
	"strings"
	"context"
	"github.com/jackc/pgx/v5"
    
)

type InsertOption struct {
	ReturnAndUpdateVals bool
	OnConflictStatement string
}


/********************************
 * Test autoincrement: {ID id uint64 false false [] [] 0 0 false [primary]}
 * 	ID uint64 id  <primary>   <autoincrement> 
 * 	IntA int int_a  
 * 	IntB SomeInts int_b  
 * 	StrA string str_a  
 * 	TimeA time.Time time_a  
 * Primary index: primary
 * Unique indexes: 
 *  primary: id
 *  test_str_a: str_a
 * Non unique indexes: 
 *  test_int_a_int_b: int_a, int_b
 ********************************/

type Tests []*Test

func (e *Test) Insert(ctx context.Context, insertOptions ...InsertOption) (err error) {
	var pgconn pgx.Conn = ctx.Value("pgconn").(pgx.Conn)
    var sql, returning string
    var args []any

    if e.ID == 0 {
        sql = `INSERT INTO "tests" (int_a, int_b, str_a, time_a)
        VALUES ($1, $2, $3, $4)`
        returning = ` RETURNING id`
        args = []any{ e.IntA, e.IntB, e.StrA, e.TimeA }
    } else {
        sql = `INSERT INTO "tests" (id, int_a, int_b, str_a, time_a)
        VALUES ($1, $2, $3, $4, $5)`
        args = []any{ e.ID, e.IntA, e.IntB, e.StrA, e.TimeA }
    }

    var returnAndUpdateVals bool
	for _, opt := range insertOptions {
		if opt.ReturnAndUpdateVals {
			returnAndUpdateVals = true
		}
		if opt.OnConflictStatement != "" {
			sql += " ON CONFLICT "+ opt.OnConflictStatement
		}
	}
    
    if returnAndUpdateVals {
        returning = ` RETURNING id, int_a, int_b, str_a, time_a`
    }
    if returning != "" {
        sql += returning
    }

	var rows pgx.Rows
	rows, err = pgconn.Query(ctx, sql, args...)
	defer func(){
		rows.Close()
		if err == nil {
			err = rows.Err()
		}
		if err != nil {
			err = fmt.Errorf("Insert query '%s' failed: %+v", sql, err)
		}
	}()

    if returnAndUpdateVals {
		if ! rows.Next() {
            // TODO: on conflict do nothing case
            return fmt.Errorf("Insert-query doesn't return anything, but has returning clause")
        }

        if err = rows.Scan(
			&e.ID,
			&e.IntA,
			&e.IntB,
			&e.StrA,
			&e.TimeA,
		); err != nil {
            return
        }
    } else if e.ID == 0 {
        if ! rows.Next() {
            // TODO: on conflict do nothing case
            return fmt.Errorf("Insert-query doesn't return anything, but has returning clause")
        }

        if err = rows.Scan(&e.ID); err != nil {
            return
        }
    }

	return nil
}

func (es Tests) Insert(ctx context.Context, insertOptions ...InsertOption) (err error) {
	var pgconn pgx.Conn = ctx.Value("pgconn").(pgx.Conn)
    var sql string
    var sqlRows []string
    var args []any

    if len(es) == 0 {
        return nil
    }

    if es[0].ID == 0 {
        sql = `INSERT INTO "tests" (id, int_a, int_b, str_a, time_a) VALUES `
        for i, e := range es {
            sqlRows = append(sqlRows, fmt.Sprintf(`(DEFAULT, $%d, $%d, $%d, $%d)`, i * 4 + 1, i * 4 + 2, i * 4 + 3, i * 4 + 4))
            args = append(args, e.IntA, e.IntB, e.StrA, e.TimeA)
        }
    } else {
        sql = `INSERT INTO "tests" (id, int_a, int_b, str_a, time_a) VALUES `
        for i, e := range es {
            sqlRows = append(sqlRows, fmt.Sprintf(`($%d, $%d, $%d, $%d, $%d)`, i * 5 + 1, i * 5 + 2, i * 5 + 3, i * 5 + 4, i * 5 + 5))
            args = append(args, e.ID, e.IntA, e.IntB, e.StrA, e.TimeA)
        }
    }

    sql += strings.Join(sqlRows, ", ")

	for _, opt := range insertOptions {
		if opt.ReturnAndUpdateVals {
			err = fmt.Errorf("ReturnAndUpdateVals option is not supported for multi-insert now")
            return
		}
		if opt.OnConflictStatement != "" {
			sql += " ON CONFLICT "+ opt.OnConflictStatement
		}
	}

	var rows pgx.Rows
	rows, err = pgconn.Query(ctx, sql, args...)
	rows.Close()
    if err == nil {
        err = rows.Err()
    }
    if err != nil {
        err = fmt.Errorf("Insert query '%s' failed: %+v", sql, err)
    }

	return
}

 
func (e *Test) Update(ctx context.Context) (err error) {
	var pgconn pgx.Conn = ctx.Value("pgconn").(pgx.Conn)

	sql := `UPDATE "tests" SET int_a = $1, int_b = $2, str_a = $3, time_a = $4	WHERE id = $5`
	var rows pgx.Rows
	rows, err = pgconn.Query(ctx, sql, e.IntA, e.IntB, e.StrA, e.TimeA, e.ID);
    rows.Close()
    if err == nil {
        err = rows.Err()
    }
    if err != nil {
        err = fmt.Errorf("Update query '%s' failed: %+v", sql, err)
    }

	return
}

func (e *Test) Delete(ctx context.Context) (err error) {
	var pgconn pgx.Conn = ctx.Value("pgconn").(pgx.Conn)

	sql := `DELETE FROM "tests" WHERE id = $1`
	var rows pgx.Rows
	rows, err = pgconn.Query(
		ctx,
		sql,
		e.ID, 
	);
    rows.Close()
    if err == nil {
        err = rows.Err()
    }
    if err != nil {
        err = fmt.Errorf("Delete query '%s' failed: %+v", sql, err)
    }

	return
}

func (es Tests) Delete(ctx context.Context) (err error) {
	var pgconn pgx.Conn = ctx.Value("pgconn").(pgx.Conn)

	sql := `DELETE FROM "tests" WHERE `
    rowsSql := make([]string, len(es))
    var args []any

    for i, e := range es {
        rowsSql[i] = fmt.Sprintf(`(id = $%d)`, i * 1 + 1)
        args = append(args, e.ID)
    }

    sql = sql + strings.Join(rowsSql, " OR ")
	var rows pgx.Rows
	rows, err = pgconn.Query(ctx, sql, args...);
    rows.Close()
    if err == nil {
        err = rows.Err()
    }
    if err != nil {
        err = fmt.Errorf("Delete query '%s' failed: %+v", sql, err)
    }

	return
}


func (Test) Find(ctx context.Context, condition string, values []interface{}) (entities Tests, err error) {

    return Test{}.Query(
        ctx,
        `SELECT id, int_a, int_b, str_a, time_a
	    FROM "tests"
	    WHERE ` + condition,
        values,
    )
}

func (Test) FindCh(ctx context.Context, condition string, values []interface{}, entitiesCh chan<- *Test, errCh chan<- error) {

    Test{}.QueryCh(
        ctx,
        `SELECT id, int_a, int_b, str_a, time_a
	    FROM "tests"
	    WHERE ` + condition,
        values,
        entitiesCh,
        errCh,
    )
}

func (Test) Query(ctx context.Context, sql string, values []interface{}) (entities Tests, err error) {

	var pgconn pgx.Conn = ctx.Value("pgconn").(pgx.Conn)

	var rows pgx.Rows
	rows, err = pgconn.Query(
		ctx,
		sql,
		values...
	)
	defer func(){
		rows.Close()
		if err == nil {
			err = rows.Err()
		}
		if err != nil {
            if len(sql) > 500 {
                sql = sql[:500] + "..."
            }

			err = fmt.Errorf("Query '%s' failed: %+v", sql, err)
		}
	}()

	for rows.Next() {

		e := Test{}

		if err = rows.Scan(
			&e.ID,
			&e.IntA,
			&e.IntB,
			&e.StrA,
			&e.TimeA,	
		); err != nil {
            return
        }

		entities = append(entities, &e)
	}

	return entities, nil
}

func (Test) QueryCh(ctx context.Context, sql string, values []interface{}, entitiesCh chan<- *Test, errCh chan<- error) {

	var (
		err error
		rows pgx.Rows
	)

    defer func(){
        if err != nil {
            errCh <- err
        }
        close(errCh)
        close(entitiesCh)
    }()

	var pgconn pgx.Conn = ctx.Value("pgconn").(pgx.Conn)

	rows, err = pgconn.Query(ctx, sql, values...)
	defer func(){
		rows.Close()
		if err == nil {
			err = rows.Err()
		}
		if err != nil {
			if len(sql) > 500 {
				sql = sql[:500] + "..."
			}

			err = fmt.Errorf("Query '%s' failed: %+v", sql, err)
		}
	}()

    if err != nil {
        return
    }

    for rows.Next() {

		e := Test{}

		if err = rows.Scan(
			&e.ID,
			&e.IntA,
			&e.IntB,
			&e.StrA,
			&e.TimeA,	
		); err != nil {
            errCh <- err
            return
        }

		entitiesCh <- &e
	}

    return
}

func (e Test) GetAll(ctx context.Context) (Tests, error) {
	return e.Find(ctx, "1=1", []any{})
}

func (e Test) GetAllCh(ctx context.Context, entitiesCh chan<- *Test, errCh chan<- error) {
	e.FindCh(ctx, "1=1", []any{}, entitiesCh, errCh)
}

 
func (e Test) GetByPrimary(ctx context.Context, id uint64) (*Test, error) {
	es, err := e.Find(
		ctx,
		"id = $1",
		[]any{ id },
	)
	if err != nil {
		return nil, err
	}
	if len(es) == 1 {
		return es[0], nil
	}

	return nil, nil
}

func (e Test) MultiGetByPrimary(ctx context.Context, id []uint64) (Tests, error) {
	
	var params []any = make([]any, 0, len(id) * 1)

	where := make([]string, len(id))
	for i := range id {
		where[i] = fmt.Sprintf("(id = $%d)", 1 + i)
		params = append(params, id[i])
	}

	return e.Find(ctx, strings.Join(where, " OR "), params)
}

func (e Test) MultiGetByPrimaryCh(ctx context.Context, id []uint64, entitiesCh chan<- *Test, errCh chan<- error) {
	
	var params []any = make([]any, 0, len(id) * 1)

	where := make([]string, len(id))
	for i := range id {
		where[i] = fmt.Sprintf("(id = $%d)", 1 + i)
		params = append(params, id[i])
	}

	e.FindCh(ctx, strings.Join(where, " OR "), params, entitiesCh, errCh)
}
func (e Test) GetByTestStrA(ctx context.Context, strA string) (*Test, error) {
	es, err := e.Find(
		ctx,
		"str_a = $1",
		[]any{ strA },
	)
	if err != nil {
		return nil, err
	}
	if len(es) == 1 {
		return es[0], nil
	}

	return nil, nil
}

func (e Test) MultiGetByTestStrA(ctx context.Context, strA []string) (Tests, error) {
	
	var params []any = make([]any, 0, len(strA) * 1)

	where := make([]string, len(strA))
	for i := range strA {
		where[i] = fmt.Sprintf("(str_a = $%d)", 1 + i)
		params = append(params, strA[i])
	}

	return e.Find(ctx, strings.Join(where, " OR "), params)
}

func (e Test) MultiGetByTestStrACh(ctx context.Context, strA []string, entitiesCh chan<- *Test, errCh chan<- error) {
	
	var params []any = make([]any, 0, len(strA) * 1)

	where := make([]string, len(strA))
	for i := range strA {
		where[i] = fmt.Sprintf("(str_a = $%d)", 1 + i)
		params = append(params, strA[i])
	}

	e.FindCh(ctx, strings.Join(where, " OR "), params, entitiesCh, errCh)
}

 
func (e Test) GetByTestIntAIntB(ctx context.Context, intA int, intB SomeInts) (Tests, error) {
	return e.Find(
		ctx,
		"int_a = $1 AND int_b = $2",
		[]any{ intA, intB },
	)
}

func (e Test) GetByTestIntAIntBCh(ctx context.Context, intA int, intB SomeInts, entitiesCh chan<- *Test, errCh chan<- error) {
	e.FindCh(
		ctx,
		"int_a = $1 AND int_b = $2",
		[]any{ intA, intB },
		entitiesCh,
		errCh,
	)
}


