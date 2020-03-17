package leetcode

/*An animal shelter, which holds only dogs and cats, operates on a strictly"first in, first out" basis. People must adopt either the"oldest" (based on arrival time) of all animals at the shelter, or they can select whether they would prefer a dog or a cat (and will receive the oldest animal of that type). They cannot select which specific animal they would like. Create the data structures to maintain this system and implement operations such as enqueue, dequeueAny, dequeueDog, and dequeueCat. You may use the built-in Linked list data structure.

enqueue method has a animal parameter, animal[0] represents the number of the animal, animal[1] represents the type of the animal, 0 for cat and 1 for dog.

dequeue* method returns [animal number, animal type], if there's no animal that can be adopted, return [-1, -1].

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/animal-shelter-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。*/

type Animal struct {
	number int
	total  int64
}
type AnimalShelf struct {
	dog []Animal
	cat []Animal
	cnt int64
}

func Constructor() AnimalShelf {
	return AnimalShelf{
		make([]Animal, 0),
		make([]Animal, 0),
		0,
	}
}

func (this *AnimalShelf) Enqueue(animal []int) {
	this.cnt++
	if animal[1] == 0 {
		this.cat = append(this.cat, Animal{animal[0], this.cnt})
	} else if animal[1] == 1 {
		this.dog = append(this.dog, Animal{animal[0], this.cnt})
	}
}

func (this *AnimalShelf) DequeueAny() []int {
	res := []int{-1, -1}
	if len(this.dog) > 0 && len(this.cat) > 0 {
		if this.dog[0].total < this.cat[0].total {
			res = []int{this.dog[0].number, 1}
			this.dog = this.dog[1:]
		} else {
			res = []int{this.cat[0].number, 0}
			this.cat = this.cat[1:]
		}
	} else if len(this.dog) > 0 {
		res = []int{this.dog[0].number, 1}
		this.dog = this.dog[1:]
	} else if len(this.cat) > 0 {
		res = []int{this.cat[0].number, 0}
		this.cat = this.cat[1:]
	}
	return res
}

func (this *AnimalShelf) DequeueDog() []int {
	res := []int{-1, -1}
	if len(this.dog) > 0 {
		res = []int{this.dog[0].number, 1}
		this.dog = this.dog[1:]
	}
	return res
}

func (this *AnimalShelf) DequeueCat() []int {
	res := []int{-1, -1}
	if len(this.cat) > 0 {
		res = []int{this.cat[0].number, 0}
		this.cat = this.cat[1:]
	}
	return res
}

/**
 * Your AnimalShelf object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Enqueue(animal);
 * param_2 := obj.DequeueAny();
 * param_3 := obj.DequeueDog();
 * param_4 := obj.DequeueCat();
 */
