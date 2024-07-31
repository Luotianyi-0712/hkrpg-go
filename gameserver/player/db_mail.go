package player

import (
	"github.com/gucooing/hkrpg-go/pkg/database"
	"github.com/gucooing/hkrpg-go/protocol/proto"
	spb "github.com/gucooing/hkrpg-go/protocol/server"
)

func NewMail() *spb.Mail {
	return &spb.Mail{
		MailList: make(map[uint32]*spb.MailDts),
	}
}

func (g *GamePlayer) GetMail() *spb.Mail {
	db := g.GetBasicBin()
	if db.Mail == nil {
		db.Mail = NewMail()
	}
	return db.Mail
}

func NewMailDts(id uint32) *spb.MailDts {
	return &spb.MailDts{
		MailId: id,
		IsDel:  false,
		IsRead: false,
	}
}

func (g *GamePlayer) GetMailList() map[uint32]*spb.MailDts {
	db := g.GetMail()
	if db.MailList == nil {
		db.MailList = make(map[uint32]*spb.MailDts)
	}
	return db.MailList
}

func (g *GamePlayer) GetMailById(id uint32) *spb.MailDts {
	db := g.GetMailList()
	if db[id] == nil {
		db[id] = NewMailDts(id)
	}
	return db[id]
}

func (g *GamePlayer) ReadMail(id uint32) {
	if id == 0 {
		return
	}
	db := g.GetMailById(id)
	db.IsRead = true
}

func (g *GamePlayer) DelMail(id uint32) {
	if id == 0 {
		return
	}
	db := g.GetMailById(id)
	db.IsDel = true
}

// TODO 邮件奖励兑换方法（拓展此处以支持更多奖励物品
func (g *GamePlayer) MailReadItem(itemList []*database.Item) bool {
	allSync := &AllPlayerSync{
		IsBasic:      true,
		MaterialList: make([]uint32, 0),
		AvatarList:   make([]uint32, 0),
	}
	pileItem := make([]*Material, 0)
	for _, item := range itemList {
		switch item.ItemType {
		case database.MailAvatar:
			allSync.AvatarList = append(allSync.AvatarList, item.ItemId)
			g.AddAvatar(item.ItemId, proto.AddAvatarSrcState_ADD_AVATAR_SRC_NONE)
		case database.MailMaterial:
			allSync.MaterialList = append(allSync.MaterialList, item.ItemId)
			pileItem = append(pileItem, &Material{
				Tid: item.ItemId,
				Num: item.Num,
			})
		}
	}
	g.AddMaterial(pileItem)
	g.AllPlayerSyncScNotify(allSync)
	return true
}

func (g *GamePlayer) GetAllMail() []*proto.ClientMail {
	mailList := make([]*proto.ClientMail, 0)
	mailMap := database.GetAllMail()
	for id, mail := range mailMap {
		db := g.GetMailById(id)
		if db.IsDel {
			continue
		}
		pbMail := &proto.ClientMail{
			IsRead:     db.IsRead,
			ExpireTime: mail.EndTime.Time.Unix(),
			Time:       mail.BeginTime.Time.Unix(),
			TemplateId: 0,
			Attachment: &proto.ItemList{ // 奖励
				ItemList: g.GetAttachment(mail.ItemList),
			},
			Title:       mail.Title,
			Sender:      mail.Sender,
			ParaList:    nil, // 参数
			Id:          mail.Id,
			MessageText: mail.Content,
		}
		mailList = append(mailList, pbMail)
	}
	return mailList
}

func (g *GamePlayer) GetAttachment(itemList []*database.Item) []*proto.Item {
	ItemList := make([]*proto.Item, 0)
	for _, item := range itemList {
		Item := &proto.Item{
			ItemId: item.ItemId,
			Num:    item.Num,
		}
		ItemList = append(ItemList, Item)
	}
	return ItemList
}
