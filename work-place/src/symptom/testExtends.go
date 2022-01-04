package symptom

type Student struct {

	name string
	age int
	Address string
}

type Seion struct {

	habbit string
	Student
}

type Collage struct {

	job string
	Student
}

func (s *Student) GetName() string{

	return s.name
}

func (s *Student) SetName(name string) {

	s.name = name
}

func (high *Seion) Gethabbit() string {

   return high.habbit
}

func (high *Seion) SetHabbit(habbit string) {

	high.habbit = habbit
}

func (c *Collage) GetJob() string {

	return c.job
}

func (c *Collage) SetJob(job string) {

	c.job = job
}