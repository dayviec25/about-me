package data

import "about-me/model"

func GetAboutMeData() (model.AboutMe, error) {
	Reviews := []model.Review{
		{
			Author:  "Tim Jensen",
			Company: "FantasyPros",
			Quote:   "Working with David at FantasyPros was an exceptional experience. His expertise in mobile app development, especially in the Kotlin environment, significantly enhanced our product offerings. David's innovative approach to problem-solving and his commitment to quality have left a lasting impact on our team.",
		},
		{
			Author:  "Rishi Tapsee",
			Company: "Yellow Pages",
			Quote:   "David's tenure at Yellow Pages was marked by his remarkable skill in API development and data integration. His ability to architect robust backend solutions played a crucial role in our project's success. David's blend of technical proficiency and collaborative spirit makes him an invaluable asset.",
		},
		{
			Author:  "Sanjay Thakur",
			Company: "Verdict MMA",
			Quote:   "David's role at Verdict MMA was transformative. As a founder, he not only demonstrated outstanding technical leadership but also showed a keen understanding of the MMA market. His strategic vision and drive for innovation were pivotal in scaling our platform.",
		},
		{
			Author:  "Sucheta Goregta",
			Company: "Cox Automotive",
			Quote:   "David's expertise at Cox Automotive was a game-changer, especially his work with ETL services and financial software solutions for the automotive industry. His attention to detail and ability to manage complex datasets effectively streamlined our operations and enhanced our service quality.",
		},
	}
	Traits := []string{
		"Mobile Engineer", "Innovator", "Self-Starter", "Quick Learner", "Passionate", "Creative", "Team Player", "Problem Solver", "Leader", "Mentor", "Teacher", "Learner", "Instructor", "Speaker", "Writer",
	}
	WorkExperience := []model.WorkExperience{
		{
			Date:        "2021 - Present",
			Company:     "Verdict MMA",
			Role:        "Full Stack Mobile Engineer",
			Description: "Founder of Verdict MMA, a groundbreaking MMA social platform with an integrated sports betting feature. Successfully led the company from its seed funding stage, overseeing the development of a robust tech stack and the introduction of innovative features. Pioneered process enhancements to streamline code deployment, significantly improving developer efficiency and product quality",
		},
		{
			Date:        "2018 - 2021",
			Company:     "FantasyPros",
			Role:        "Senior Android Engineer",
			Description: "Led the creation and maintenance of 6 Android apps at FantasyPros, catering to fantasy sports enthusiasts. Achieved a notable increase in average app ratings from 4.1 to 4.7, and doubled the Monthly Recurring Revenue (MMR), positioning these apps among the top 10 in the sports category. Successfully transitioned the entire codebase to Kotlin, enhancing app performance and maintainability. Developed specialized Android tools for major sports brands including ESPN, Yahoo, The Athletic, and NBA.",
		},
		{
			Date:        "2016 - 2018",
			Company:     "Cox Automotive",
			Role:        "Full Stack Engineer",
			Description: "Created ETL services daily to gather data on all financial lease and loan offers in the country for all cars. Create financial software for dealership including Mazda, Volvo, Hyundai, Porsche.",
		},
		{
			Date:        "2014 - 2016",
			Company:     "Yellow Pages",
			Role:        "Backend Engineer",
			Description: "As a Backend Engineer at Yellow Pages, I played a key role in developing and maintaining the API for the Yellow Pages mobile app. I engineered a secure, RESTful API using .NET, specifically designed to handle spatial data for real estate listings. My responsibilities also included deploying an ETL solution to seamlessly integrate over 200,000 resale and rental listings daily, as well as automating services for email templating and daily statistical reporting.",
		},
	}
	Projects := []model.Project{
		{
			Title:        "Verdict MMA Picks & Scoring",
			Logo:         "https://play-lh.googleusercontent.com/w7s5HrUc0JgcR_vYRtd-PlrnVJ565hDhkK2V8hVcOf2dWWaPm5rCVCVrQOo_LsI1-eY=w240-h480-rw",
			LogoAlt:      "Verdict MMA Logo",
			ReviewRating: "4.7",
			ReviewCount:  "977",
			Downloads:    "50k+",
			CompanyName:  "Verdict MMA",
			Link:         model.Link{Href: "https://play.google.com/store/apps/details?id=com.verdictmma.verdict&hl=en_CA&gl=US", Label: "Visit Verdict MMA"},
		},
		{
			Title:        "Fantasy Football My Playbook",
			Logo:         "https://play-lh.googleusercontent.com/w9THLA8cUKvIRHgH_bhTJsjTk3hYl81GOI6fZkB3dAI1I_6-o4LxV0CYLl65bqkce4oT=w240-h480-rw",
			ReviewRating: "4.9",
			ReviewCount:  "26.7k",
			Downloads:    "100k+",
			LogoAlt:      "MyPlaybook Logo",
			CompanyName:  "FantasyPros",
			Link:         model.Link{Href: "https://play.google.com/store/apps/details?id=com.fantasypros.myplaybook&hl=en_CA&gl=US", Label: "Visit Fantasy Football My Playbook"},
		},
		{
			Title:        "Fantasy Football Draft Wizard",
			Logo:         "https://play-lh.googleusercontent.com/MNMUUUznoxez9vBomF1O_DLY_n39WQ2nojWTc3X28rPV21PxIt78k9yWQhFeG5PBgg=w240-h480-rw",
			ReviewRating: "4.5",
			ReviewCount:  "12.2k",
			Downloads:    "500k+",
			LogoAlt:      "Draft Wizard Logo",
			CompanyName:  "FantasyPros",
			Link:         model.Link{Href: "https://play.google.com/store/apps/details?id=com.fantasypros.draftwizard.football&hl=en_CA&gl=US", Label: "Visit Fantasy Football Draft Wizard"},
		},
		{
			Title:        "Fantasy News & Scores",
			Logo:         "https://play-lh.googleusercontent.com/HreX3kivqhUegukOEyT2MQ9_9S2Gb50B9KzT6qlEaf5eBDJGSfgEmsAGwRqNakU6GKg=w240-h480-rw",
			ReviewRating: "4.9",
			ReviewCount:  "5.57k",
			Downloads:    "100k+",
			LogoAlt:      "Fantasy News Logo",
			CompanyName:  "FantasyPros",
			Link:         model.Link{Href: "https://play.google.com/store/apps/details?id=com.fantasypros.news&hl=en_CA&gl=US", Label: "Visit Fantasy News"},
		},
	}
	Contact := model.Contact{
		Name:              "David Chung",
		ProfilePictureURL: "https://pbs.twimg.com/profile_images/1713322706609827840/r7lsuZUg_400x400.jpg",
		CallToActionLink:  "https://calendly.com/dvlchung/1-hour-meeting",
		GithubLink:        "https://www.github.com/dchungdev",
		LinkedInLink:      "https://www.linkedin.com/in/dchungdev/",
		TwitterLink:       "https://www.twitter.com/dayviec",
		EmailLink:         "mailto:dvlchung@gmail.com",
	}
	Blogs := []model.Blog{
		{
			Title:   "How to Build a Blog with Golang and DigitalOcean Spaces",
			Content: "This is a tutorial on how to build a blog with Golang and DigitalOcean Spaces.",
			Date:    "2021-08-01",
			Author:  "David Chung",
		},
	}

	return model.AboutMe{
		Title:          "Full Stack Mobile Engineer",
		Description:    "A seasoned Full Stack Mobile Engineer with over 10 years of experience, specializing in developing, modernizing, and maintaining cutting-edge mobile applications and web platforms. Known for my ability to innovate and lead in fast-paced environments, I bring a proven track record of enhancing user experience, streamlining development processes, and forging strategic partnerships across various industries.",
		WorkExperience: WorkExperience,
		Projects:       Projects,
		Reviews:        Reviews,
		Traits:         Traits,
		Blogs:          Blogs,
		Contact:        Contact,
	}, nil
}
