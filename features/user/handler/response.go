package handler

import "belajar/bareng/features"

type Response struct {
	Id    uint   `json:"id"`
	Nama  string `json:"nama,omitempty"`
	NoTlp string `json:"no_tlp,omitempty"`
	Email string `json:"email,omitempty"`
	Alamat string `json:"alamat,omitempty"`
}

func EntityToResponse(input features.UserEntity) Response{
	return Response{
		Id:		input.Id,
		Nama: 	input.Nama,
		NoTlp: 	input.NoTlp,
		Email: 	input.Email,
		Alamat: input.Alamat,
	}
}