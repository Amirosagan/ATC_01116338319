package seed

import (
	"log"
	"time"

	"online-task/internal/models"
	"online-task/pkg/database"

	"github.com/google/uuid"
)

var techEvents = []struct {
	Name        string
	Description string
	Category    string
	Date        time.Time
	Location    string
	Price       float64
	Image       string
}{
	{
		Name:        "RSA Conference 2025",
		Description: "Join the world's leading cybersecurity professionals for four days of cutting-edge insights, networking, and hands-on experience. Explore the latest trends in cybersecurity, from AI-driven security solutions to zero-trust architectures.",
		Category:    "Cybersecurity",
		Date:        time.Date(2025, time.April, 28, 9, 0, 0, 0, time.UTC),
		Location:    "Moscone Center, San Francisco, CA",
		Price:       2499.99,
		Image:       "https://live.staticflickr.com/65535/54488513037_58956cbf05_c.jpg",
	},
	{
		Name:        "AWS re:Invent 2024",
		Description: "The largest gathering of the global cloud computing community. Join us for 5 days of keynotes, technical sessions, workshops, and networking opportunities focused on AWS cloud solutions.",
		Category:    "Cloud Computing",
		Date:        time.Date(2024, time.December, 2, 9, 0, 0, 0, time.UTC),
		Location:    "Las Vegas, Nevada",
		Price:       1799.99,
		Image:       "https://media.gettyimages.com/id/1245247600/photo/invent-2022.jpg?s=2048x2048&w=gi&k=20&c=Hh_S13DUI2pa44oh_HC0tFYGloVomuVDajB9ADI15NI=",
	},
	{
		Name:        "CES 2025",
		Description: "Experience the world's most influential tech event, showcasing breakthrough technologies and global innovators. From AI and robotics to IoT and digital health, witness the future of technology.",
		Category:    "Technology",
		Date:        time.Date(2025, time.January, 7, 10, 0, 0, 0, time.UTC),
		Location:    "Las Vegas Convention Center, NV",
		Price:       999.99,
		Image:       "https://techcrunch.com/wp-content/uploads/2024/01/GettyImages-1246002029.jpg?resize=1200,800",
	},
	{
		Name:        "Dublin Tech Summit 2025",
		Description: "Join 8,000+ attendees, 200+ speakers, and 150+ startups at Europe's fastest-growing tech conference. Explore AI, blockchain, cybersecurity, and more while networking with industry leaders.",
		Category:    "Technology",
		Date:        time.Date(2025, time.May, 29, 9, 0, 0, 0, time.UTC),
		Location:    "RDS, Dublin, Ireland",
		Price:       599.99,
		Image:       "https://dublintechweek.com/wp-content/uploads/2024/12/DTS25-2.jpg",
	},
	{
		Name:        "AI & Big Data Expo 2024",
		Description: "Discover the latest innovations in AI and Big Data with 7,000+ attendees and 250+ speakers. Features 18 conference tracks covering machine learning, data analytics, and AI implementation strategies.",
		Category:    "AI & Data Science",
		Date:        time.Date(2024, time.June, 5, 9, 0, 0, 0, time.UTC),
		Location:    "Santa Clara Convention Center, CA",
		Price:       899.99,
		Image:       "https://www.ai-expo.net/wp-content/uploads/2023/07/The-Future-of-enterprise-technology-24-1024x682-1.webp",
	},
}

func SeedEvents() error {
	db := database.GetDB()

	// Check if events already exist
	var count int64
	db.Model(&models.Event{}).Count(&count)
	if count > 0 {
		log.Println("Events already seeded")
		return nil
	}

	// Create events
	for _, eventData := range techEvents {
		event := models.Event{
			ID:          uuid.New().String(),
			Name:        eventData.Name,
			Description: eventData.Description,
			Category:    eventData.Category,
			Date:        eventData.Date,
			Location:    eventData.Location,
			Price:       eventData.Price,
			Image:       eventData.Image,
		}

		if err := db.Create(&event).Error; err != nil {
			return err
		}
	}

	log.Printf("Successfully seeded %d events\n", len(techEvents))
	return nil
}
