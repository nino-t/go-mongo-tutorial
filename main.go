package main

import (
	"fmt"
	"time"

	"github.com/nino-t/go-mongo-tutorial/config"
	"github.com/nino-t/go-mongo-tutorial/src/modules/profile/model"
	"github.com/nino-t/go-mongo-tutorial/src/modules/profile/repository"
)

func main() {
	fmt.Println("Go + MongoDB is running...")

	db, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	}

	profileRepository := repository.NewProfileRepositoryMongo(db, "profile")
	// saveProfile(profileRepository)
	// updateProfile(profileRepository)
	// deleteProfile(profileRepository)
	// getProfile("U1", profileRepository)
	getProfiles(profileRepository)
}

func saveProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile
	p.ID = "U1"
	p.FirstName = "Trisno"
	p.LastName = "Nino"
	p.Email = "ninotrisno34@gmail.com"
	p.Password = "123456"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Save(&p)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile saved...")
	}
}

func updateProfile(profileRepository repository.ProfileRepository) {
	var p model.Profile
	p.ID = "U1"
	p.FirstName = "Trisno"
	p.LastName = "Nino"
	p.Email = "ninotrisno34@gmail.com"
	p.Password = "admin123"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := profileRepository.Update("U1", &p)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile updated...")
	}
}

func deleteProfile(profileRepository repository.ProfileRepository) {
	err := profileRepository.Delete("U1")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile deleted...")
	}
}

func getProfile(id string, profileRepository repository.ProfileRepository) {
	profile, err := profileRepository.FindByID(id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile fetch one...")
		fmt.Println(profile.ID)
		fmt.Println(profile.FirstName)
		fmt.Println(profile.LastName)
		fmt.Println(profile.Email)
	}
}

func getProfiles(profileRepository repository.ProfileRepository) {
	profiles, err := profileRepository.FindByAll()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Profile fetch all...")

		for _, profile := range profiles {
			fmt.Println(profile.ID)
			fmt.Println(profile.FirstName)
			fmt.Println(profile.LastName)
			fmt.Println(profile.Email)
			fmt.Println("----")
		}
	}
}
