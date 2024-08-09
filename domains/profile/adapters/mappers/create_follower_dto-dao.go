package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/profile/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/intermediate_tables/follower/dao"
)

func CreateFollowerDTOtoDAO(followerDTO *dto.CreateFollowerDTOReq, convert utils.ConvertToObjectIDFunc) (*dao.CreateFollowerDAOReq, error) {
	fromUseroOID, err := convert(followerDTO.From)
	if err != nil {
		return nil, err
	}

	toUseroOID, err := convert(*followerDTO.ToUser)
	if err != nil {
		return nil, err
	}

	followerDAO := &dao.CreateFollowerDAOReq{
		From:   *fromUseroOID,
		ToUser: toUseroOID,
	}

	return followerDAO, nil
}
