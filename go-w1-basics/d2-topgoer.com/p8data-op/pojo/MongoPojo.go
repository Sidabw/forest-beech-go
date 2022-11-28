package pojo

type RoomSetting struct {
	Id         string `bson:"_id,omitempty"`
	Menu       string `bson:"menu,omitempty"`
	PicTxtMenu string `bson:"picTxtMenu,omitempty"`
	RoomId     int    `bson:"roomId"`
}
