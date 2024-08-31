package mappers

import (
	"github.com/DBrange/didis-comp-bk/cmd/api/utils"
	"github.com/DBrange/didis-comp-bk/domains/chat/models/dto"
	"github.com/DBrange/didis-comp-bk/domains/repository/models/chat_message/dao"
)

func CreateChatMessageDTOtoDAO(chatMessageDTO *dto.CreateChatMessageDTOReq, convert utils.ConvertToObjectIDFunc) (*dao.CreateChatMessageDAOReq, error) {
	chatOID, err := convert(chatMessageDTO.ChatID)
	if err != nil {
		return nil, err
	}
	senderOID, err := convert(chatMessageDTO.SenderID)
	if err != nil {
		return nil, err
	}

	chatMessageDAO := &dao.CreateChatMessageDAOReq{
		ChatID:   chatOID,
		SenderID: senderOID,
		Content:  chatMessageDTO.Content,
	}

	return chatMessageDAO, nil
}
