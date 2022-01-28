package database

import (
	"context"
	"log"
	"time"

	model "github.com/Peyton232/dual-go-reciever/models"
	"go.mongodb.org/mongo-driver/bson"
)

type metaData struct {
	Name        string
	Description string `json:"description"`
	Image       string `json:"image"`
	Attributes  Attr   `json:"attributes"`
}

type Attr struct {
	DisplayType string `json:"display_type"`
	TraitType   string `json:"trait_type"`
	Value       string `json:"value"`
}

//Creation
func (db *DB) ImageDecode() bool {
	collection := db.dataIrid
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Print(err)
		log.Print("\nunable to find any data from DB in database package\n")
		return false
	}
	var data []*model.SpaceData
	var lastDat string = "randomString"
	for cur.Next(ctx) {
		var curDat *model.SpaceData
		err := cur.Decode(&curDat)
		if err != nil {
			log.Print(err)
			log.Print("\nunable to read user model in database package\n")
			return false
		}
		if *curDat.Data == lastDat {
			//filter out duplicates
		} else {
			lastDat = *curDat.Data
			data = append(data, curDat)
		}

	}

	// var rawImgData string
	for _, v := range data {
		byteSlice := []byte(*v.Data)
		print(string(byteSlice[:len(byteSlice)]))
		print("   ")
		print(string(byteSlice[len(byteSlice)-1]))
		print("\n\n")
		// print(*v.Data)
	}

	// print("\n\n" + rawImgData + "\n\n")

	// //remove spaces
	// rawImgData = strings.ReplaceAll(rawImgData, " ", "")

	// bs, err := hex.DecodeString(rawImgData)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(bs))

	// imgbytes, _ := base64.StdEncoding.DecodeString(string(bs))
	// jpgI, errJpg := jpeg.Decode(bytes.NewReader(imgbytes))
	// if errJpg != nil {
	// 	log.Fatal(errJpg)
	// }

	// f, err := os.Create("outimage.jpg")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	// // Encode to `PNG` with `DefaultCompression` level
	// // then save to file
	// err = jpeg.Encode(f, jpgI, &jpeg.Options{Quality: 90})
	// if err != nil {
	// 	log.Fatal("Image encode failed")
	// }

	// nft := metaData{}
	// nft.Image = "data:image/jpg;base64," + string(bs)
	// nft.Name = "name"
	// nft.Description = "some description"
	// nft.Attributes.DisplayType = "display type"
	// nft.Attributes.TraitType = "trait type"
	// nft.Attributes.Value = "value"

	// b, err := json.Marshal(nft)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return false
	// }
	// fmt.Println(string(b))

	return true
}
