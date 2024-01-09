package room

type CreateRoomRequest struct {
	Name string `json:"name"`
}

type AddMemberRequest struct {
	MemberId uint `json:"userId"`
}

type RemoveMemberRequest struct {
	MemberId uint `json:"userId"`
}
