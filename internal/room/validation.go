package room

type CreateRoomRequest struct {
	Name string `json:"name"`
}

type AddMemberRequest struct {
	MemberId string `json:"userId"`
}

type RemoveMemberRequest struct {
	MemberId string `json:"userId"`
}
