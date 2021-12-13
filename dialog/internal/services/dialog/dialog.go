package dialog

import (
	"database/sql"
	"golang.org/x/net/context"
	"time"
)

const queryTimeout = time.Second * 5

const (
	RoleCreator = "creator"
	RoleMember  = "member"
)

type Service struct {
	db *sql.DB
}

func NewDialogService(db *sql.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreateDialog(form CreateDialogForm) (Dialog, error) {
	ctx, _ := context.WithTimeout(context.Background(), queryTimeout)
	query := `insert into dialog(name, creator_id) values(?,?)`

	result, err := s.db.ExecContext(ctx, query, form.Name, form.UserId)

	if err != nil {
		return Dialog{}, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return Dialog{}, err
	}

	err = s.AddMember(AddMemberForm{
		DialogId: id,
		MemberId: form.UserId,
		Role:     RoleCreator,
	})

	if err != nil {
		return Dialog{}, err
	}

	return s.GetById(id)
}

func (s *Service) AddMember(form AddMemberForm) error {
	ctx, _ := context.WithTimeout(context.Background(), queryTimeout)
	query := `insert into dialog_member(dialog_id, member_id, role) values(?,?,?)`

	_, err := s.db.ExecContext(ctx, query, form.DialogId, form.MemberId, form.Role)

	return err
}

func (s *Service) GetById(id int64) (Dialog, error) {
	query := `select * from dialog where dialog_id=?`
	ctx, _ := context.WithTimeout(context.Background(), queryTimeout)

	row := s.db.QueryRowContext(ctx, query, id)
	d := Dialog{}
	err := row.Scan(&d.DialogId, &d.Name, &d.CreatorId, &d.CreateAt)

	if err != nil {
		return d, err
	}

	return d, err
}
