package tplfunc

func ImageUploadPath() string {
	return Pwd()+"/upload/image/" + Date("Y") + "/" + Date("m")
}