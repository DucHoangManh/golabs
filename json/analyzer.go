package main

import (
	"sort"
	"strconv"
	"strings"
	"time"
)

//2.1 Gom tất cả những người trong cùng một thành phố lại
func GroupPeopleByCity(p []Person) (result map[string][]Person) {
	result = make(map[string][]Person)
	for _, person := range p {
		result[person.City] = append(result[person.City], person)
	}
	return result
}
//2.2 Nhóm các nghề nghiệp và đếm số người làm
func GroupPeopleByJob(p []Person) (result map[string]int) {
	result = make(map[string]int)
	for _, person := range p {
		result[person.Job] ++ 
	}
	return
}
//2.3 Tìm 5 nghề có nhiều người làm nhất, đếm từ cao xuống thấp
func Top5JobsByNumer(p []Person) []string{
	peopleByJob := GroupPeopleByJob(p)
	jobs := make([]string, len(peopleByJob))
	i:=0
	for key := range peopleByJob{
		jobs[i]=key
		i++
	} 
	sort.SliceStable(jobs, func(i, j int) bool {
		return peopleByJob[jobs[i]] > peopleByJob[jobs[j]]
	})
	return jobs[:5]
}
//2.4 Tìm 5 thành phố có nhiều người trong danh sách ở nhất, đếm từ cao xuống thấp
func CountPeopleByCity(p []Person) (result map[string]int){
	result = make(map[string]int)
		for _, person := range p {
		result[person.City] ++ 
	}
	return
}
func Top5CitiesByNumer(p []Person) (result []string){
	peopleByCity := CountPeopleByCity(p)
	cities := make([]string, len(peopleByCity))
	i:=0
	for key := range peopleByCity{
		cities[i]=key
		i++
	} 
	sort.SliceStable(cities, func(i, j int) bool {
		return peopleByCity[cities[i]] > peopleByCity[cities[j]]
	})
	return cities[:5]
}
//2.5 Trong mỗi thành phố, hãy tìm ra nghề nào được làm nhiều nhất
//Trong trường hợp chỉ cần tìm ra một nghề được làm nhiều nhất
func TopJobByNumerInEachCity(p []Person) (result map[string]string){
	result = make(map[string]string)
	peopleByCity := GroupPeopleByCity(p)
	for city, people := range peopleByCity{
		peopleByJob := GroupPeopleByJob(people)
		max := 0
		var jobMax string
		for job, count := range peopleByJob{
			if count > max{
				jobMax = job
				max = count
			}
		}
		result[city] = jobMax
	}
	return result
}
//Trong trường hợp có nhiều nghề cùng có số người làm bằng nhau và là nhiều nhất
func TopJobsByNumerInEachCity(p []Person) (result map[string][]string){
	result = make(map[string][]string)
	peopleByCity := GroupPeopleByCity(p)
	for city, people := range peopleByCity{
		peopleByJob := GroupPeopleByJob(people)
		topJobsByCity := make([]string,0)
		max := 0
		for _, count := range peopleByJob{
			if count > max{
				max = count
			}
		}
		for job,count := range peopleByJob{
			if count == max{
				topJobsByCity = append(topJobsByCity, job)
			}
		}
		result[city] = topJobsByCity
	}
	return result
}
//2.6 Ứng với một nghề, hãy tính mức lương trung bình
func AverageSalaryByJob(p []Person) (result map[string]int){
	result = make(map[string]int)
	sumSalaryByJob := make(map[string]int)
	countPeopleByJob := GroupPeopleByJob(p)
	for _, person := range p {
		sumSalaryByJob[person.Job]+=person.Salary
	}
	for job := range sumSalaryByJob {
		result[job] = sumSalaryByJob[job] / countPeopleByJob[job]
	}
	return
}
//2.7 Năm thành phố có mức lương trung bình cao nhất
func AverageSalaryByCity(p []Person) (result map[string]int){
	result = make(map[string]int)
	sumSalaryByCity := make(map[string]int)
	countPeopleByCity := CountPeopleByCity(p)
	for _, person := range p {
		sumSalaryByCity[person.City]+=person.Salary
	}
	for job := range sumSalaryByCity {
		result[job] = sumSalaryByCity[job] / countPeopleByCity[job]
	}
	return
}
func FiveCitiesHasTopAverageSalary(p []Person) []string{
	avgSalaryByCity := AverageSalaryByCity(p)
	cities := make([]string, len(avgSalaryByCity))
	index := 0
	for city := range avgSalaryByCity{
		cities[index] = city
		index++
	}
	sort.SliceStable(cities,func(i, j int) bool{
		return avgSalaryByCity[cities[i]]>avgSalaryByCity[cities[j]]
	})
	return cities[:5]
}
//2.8 Năm thành phố có mức lương trung bình của developer cao nhất
func FiveCitiesHasTopSalaryForDeveloper(p []Person) []string{
	peopleByCity := GroupPeopleByCity(p)
	avgDeveloperSalaryByCity := make(map[string]int)
	for city, people := range peopleByCity{
		sumDeveloperSalaryByCity := 0
		countDeveloperByCity := 0
		for _ , person := range people{
			if person.Job == "developer"{
				sumDeveloperSalaryByCity += person.Salary
				countDeveloperByCity ++
			}
		}
		if countDeveloperByCity == 0{
			avgDeveloperSalaryByCity[city] = 0
		}else {
			avgDeveloperSalaryByCity[city] = sumDeveloperSalaryByCity / countDeveloperByCity
		}
	}
	cities := make([]string,len(avgDeveloperSalaryByCity))
	index := 0
	for city := range avgDeveloperSalaryByCity{
		cities[index]=city
		index++
	}
	sort.SliceStable(cities, func(i, j int) bool {
		return avgDeveloperSalaryByCity[cities[i]] > avgDeveloperSalaryByCity[cities[j]]
	})
	return cities[:5]
}
//2.9 Tuổi trung bình từng nghề nghiệp
func getAge(birthdayString string) int {
	birthdaySlice := strings.Split(birthdayString, "-")
	by, _ :=strconv.Atoi(birthdaySlice[0])
	bm, _ :=strconv.Atoi(birthdaySlice[1])
	bd, _ :=strconv.Atoi(birthdaySlice[2])
	birthday := time.Date(by, time.Month(bm), bd, 0, 0, 0, 0, time.UTC)
	now := time.Now()	
	years := now.Year() - birthday.Year()
	if now.YearDay() < birthday.YearDay(){
		years--
	}
	return years
}
func AverageAgePerJob(p []Person) (result map[string]float32){
	result = make(map[string]float32)
	peopleNumberByJob := GroupPeopleByJob(p)
	sumOfPeopleAgeByJob := make(map[string]int)
	for _, person := range p{
		sumOfPeopleAgeByJob[person.Job]+=getAge(person.Birthday)
	}
	for job, sumAge := range sumOfPeopleAgeByJob{
		result[job] = float32(sumAge) / float32(peopleNumberByJob[job]) 
	}
	return 
}
//2.10 Tuổi trung bình ở từng thành phố
func AverageAgePerCity(p []Person) (result map[string]float32){
	result = make(map[string]float32)
	peopleNumberByCity := CountPeopleByCity(p)
	sumOfPeopleAgeByCity := make(map[string]int)
	for _, person := range p{
		sumOfPeopleAgeByCity[person.City]+=getAge(person.Birthday)
	}
	for city, sumAge := range sumOfPeopleAgeByCity{
		result[city] = float32(sumAge) / float32(peopleNumberByCity[city]) 
	}
	return
}