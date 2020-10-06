package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	var a int
	var b int
	var key = 1
	var key2 = 1
	var key3 = 1
	var mon int
	var hp1 = 100
	var hp2 = 75
	var hp2or = hp2
	var hp1or = hp1
	var attack1 = 10
	var attack2 = 30
	var attack3 = 20
	var attack4 = 50
	var attacken1 = 10
	var attacken2 = 30 
	var attacken3 = 10

	fmt.Println("You wake up, lost, in a forest, with no memory to recall.")
	time.Sleep(4 * time.Second)
	fmt.Println("The leaves of the trees, blocked the sunlight from the two suns.")
	time.Sleep(4 * time.Second)
	fmt.Println("You notice a sword in your hand, in horrible condition.") 
	time.Sleep(4 * time.Second)
	fmt.Println("Confused, you get up, despite the physical pain, You march forward in hope to find life, but it is hopeless.")
	time.Sleep(4 * time.Second)
	fmt.Println("Suddenly, you are shocked with an impact to the back of your head. It's a Wood Dweller.")
	time.Sleep(4 * time.Second)
	fmt.Println("you enter a battle with a Wood Dweller")
	for battle1 := 1 ;; battle1++ {
		if battle1 % 10 == 0 {
			fmt.Println("ultimate attack is ready")
			time.Sleep(1 * time.Second)
			fmt.Println("would you like to use ultimate attack")
			fmt.Println("1.Yes")
			fmt.Println("2.no")
			fmt.Scanln(&b)
			if b == 1 {
				time.Sleep(1 * time.Second)
				fmt.Println("using ultimate attack")
				time.Sleep(5 * time.Second)
				hp2 = hp2 - attack4
				fmt.Println("enemy took",attack4,"damage")
				if hp2 <= 0 {
					fmt.Println("You Win")
					break
				}
			} else if b == 2 {
				time.Sleep(1 * time.Second)
				continue
			} else {
				time.Sleep(1 * time.Second)
				fmt.Println("sorry that is an invalid command")
				time.Sleep(1 * time.Second)
				continue
			}
		}
		fmt.Println("your hp",hp1,"          enemy hp",hp2)
		fmt.Println("choose your attack")
		fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
		fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
		fmt.Println("3.heal",attack3,"health 100% accuracy")
		fmt.Scanln(&a)
		if a == 1 {
			time.Sleep(1 * time.Second)
			fmt.Println("Wood Dweller took",attack1,"damage")
			hp2 = hp2-attack1
			time.Sleep(1 * time.Second)
			if hp2 <= 0 {
				fmt.Println("You Win")
				break
			}
		} else if a == 2 {
			var s = time.Now().UnixNano()
			rand.Seed(s)
			var random = rand.Intn(3)
			if random == 1 || random == 2 || random == 3 {
				time.Sleep(1 * time.Second)
				hp2 = hp2 - attack2
			fmt.Println("Wood Dweller took",attack2,"damage")
			time.Sleep(1 * time.Second)
			if hp2 <= 0 {
				fmt.Println("You Win")
				break
			}
			} else {
				time.Sleep(1 * time.Second)
				fmt.Println("attack missed")
				time.Sleep(1 * time.Second)
			}
		} else if a == 3 {
			if hp1 > hp1or - attack3{
				time.Sleep(1 * time.Second)
				fmt.Println("You're not low enough to heal")
				time.Sleep(1 * time.Second)
			} else {
				time.Sleep(1 * time.Second)
				hp1 = hp1 + attack3
				fmt.Println("You have been healed",attack3,"health")
				time.Sleep(1 * time.Second)
			}
		} else {
			fmt.Println("Sorry that is an invalid command")
		}
		var a = time.Now().UnixNano()
		rand.Seed(a)
		var random2 = rand.Intn(2)
		time.Sleep(1 * time.Second)
		fmt.Println("enemy is attacking")
		time.Sleep(1 * time.Second)
		if random2 == 0 {
			fmt.Println("you took",attacken1,"damage")
			hp1 = hp1-attacken1
			time.Sleep(1 * time.Second)
			if hp1 <= 0 {
				fmt.Println("game over")
				return
			}
		}else if random2 == 1 {
			var b = time.Now().UnixNano()
			rand.Seed(b)
			var random3 = rand.Intn(5)
				if random3 == 1 || random3 == 2 || random3 == 3 {
					fmt.Println("you took",attacken2,"damage")
					hp1 = hp1-attacken2
					time.Sleep(1 * time.Second)	
					if hp1 <= 0 {
						fmt.Println("game over")
						return
					}
				} else {
					fmt.Println("attack missed")
						time.Sleep(1 * time.Second)	
						continue
				}
			} else {
				if hp2 > hp2or - attacken3 {
					time.Sleep(1 * time.Second)
					fmt.Println("attack missed")
					time.Sleep(1 * time.Second)	
				} else {
					time.Sleep(1 * time.Second)
					hp2 = hp2 + attacken3
					fmt.Println("enemy has been healed",attacken3,"health")
					time.Sleep(1 * time.Second)
				}
			}
		}
		time.Sleep(2 * time.Second)
		fmt.Println("The Wood Dweller drops dead, and you are victorious. You scavenge its body and claim its weapon")
		time.Sleep(2 * time.Second)
		attack1 = 20
		attack2 = 40
		attack4 = 60
		fmt.Println("Weapon Upgrade: Battle Axe")
		time.Sleep(4 * time.Second)
		fmt.Println("You venture deeper into the woods, in hope of finding something that may bring back your memory, but you feel a severe headache.")
		time.Sleep(7 * time.Second)
		for space := 1 ; space <=100 ; space++{
			fmt.Println()
		}
		fmt.Println("MAP 1: MYSTIC WOODS")
		for space := 1 ; space <=20 ; space++{
			fmt.Println()
		}
		time.Sleep(2 * time.Second)	
		for space := 1 ; space <=100 ; space++{
			fmt.Println()
		}
		time.Sleep(3 * time.Second)
		fmt.Println("After obtaining the Battle Axe, you look around the almost never-ending abyss of logs and leaves.")
		time.Sleep(3 * time.Second)
		fmt.Println("You encounter a lumbering creature, shaped identically to a walking, dead tree.")
		time.Sleep(3 * time.Second)
			attacken1 = 15
			attacken2 = 35
			attacken3 = 30
			hp2 = 100
			hp2or = hp2
			hp1 = 100
			fmt.Println("You enter a battle with Ancient Log")
			for battle1 := 1 ;; battle1++ {
				if battle1 % 10 == 0 {
					fmt.Println("ultimate attack is ready")
					time.Sleep(1 * time.Second)
					fmt.Println("would you like to use ultimate attack")
					fmt.Println("1.Yes")
					fmt.Println("2.no")
					fmt.Scanln(&b)
					if b == 1 {
						time.Sleep(1 * time.Second)
						fmt.Println("using ultimate attack")
						time.Sleep(5 * time.Second)
						hp2 = hp2 - attack4
						fmt.Println("enemy took",attack4,"damage")
						if hp2 <= 0 {
							fmt.Println("You Win")
							break
						}
					} else if b == 2 {
						time.Sleep(1 * time.Second)
						continue
					} else {
						time.Sleep(1 * time.Second)
						fmt.Println("sorry that is an invalid command")
						time.Sleep(1 * time.Second)
						continue
					}
				}
				fmt.Println("your hp",hp1,"          enemy hp",hp2)
				fmt.Println("choose your attack")
				fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
				fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
				fmt.Println("3.heal",attack3,"health 100% accuracy")
				fmt.Scanln(&a)
				if a == 1 {
					time.Sleep(1 * time.Second)
					fmt.Println("enemy took",attack1,"damage")
					hp2 = hp2-attack1
					time.Sleep(1 * time.Second)
					if hp2 <= 0 {
						fmt.Println("You Win")
						break
					}
				} else if a == 2 {
					var s = time.Now().UnixNano()
					rand.Seed(s)
					var random = rand.Intn(3)
					if random == 1 || random == 2 || random == 3 {
						time.Sleep(1 * time.Second)
						hp2 = hp2 - attack2
					fmt.Println("enemy took",attack2,"damage")
					time.Sleep(1 * time.Second)
					if hp2 <= 0 {
						fmt.Println("You Win")
						break
					}
					} else {
						time.Sleep(1 * time.Second)
						fmt.Println("attack missed")
						time.Sleep(1 * time.Second)
					}
				} else if a == 3 {
					if hp1 > hp1or - attack3 {
						time.Sleep(1 * time.Second)
						fmt.Println("You're not low enough to heal")
						time.Sleep(1 * time.Second)
					} else {
						time.Sleep(1 * time.Second)
						hp1 = hp1 + attack3
						fmt.Println("You have been healed",attack3,"health")
						time.Sleep(1 * time.Second)
					}
				} else {
					fmt.Println("Sorry that is an invalid command")
				}
				var a = time.Now().UnixNano()
				rand.Seed(a)
				var random2 = rand.Intn(2)
				time.Sleep(1 * time.Second)
				fmt.Println("enemy is attacking")
				time.Sleep(1 * time.Second)
				if random2 == 0 {
					fmt.Println("you took",attacken1,"damage")
					hp1 = hp1-attacken1
					time.Sleep(1 * time.Second)
					if hp1 <= 0 {
						fmt.Println("game over")
						return
					}
				} else if random2 == 1 {
					var b = time.Now().UnixNano()
					rand.Seed(b)
					var random3 = rand.Intn(5)
						if random3 == 1 || random3 == 2 || random3 == 3 {
							fmt.Println("you took",attacken2,"damage")
							hp1 = hp1-attacken2
							time.Sleep(1 * time.Second)	
							if hp1 <= 0 {
								fmt.Println("game over")
								return
							}
						} else {
							fmt.Println("attack missed")
								time.Sleep(1 * time.Second)	
								continue
						}
					} else {
						if hp2 > hp2or - attacken3 {
							time.Sleep(1 * time.Second)
							fmt.Println("attack missed")
							time.Sleep(1 * time.Second)	
						} else {
							time.Sleep(1 * time.Second)
							hp2 = hp2 + attacken3
							fmt.Println("enemy has been healed",attacken3,"health")
							time.Sleep(1 * time.Second)
					}
				}
			}
			time.Sleep(1 * time.Second)
		fmt.Println("You gained 20 coins")
		time.Sleep(1 * time.Second)
		mon = mon + 20
		fmt.Println("Total coins:",mon)
		time.Sleep(3 * time.Second)
		fmt.Println("You emerge victorious. After the pain is relieved, You venture deep into the woods.")
		time.Sleep(3 * time.Second)
		fmt.Println("Wary of your surrounding,")
		time.Sleep(3 * time.Second)
		fmt.Println("until you notice an unusual structure off in the distance But you hear some sounds in the opposite Direction.")
		time.Sleep(3 * time.Second)
		fmt.Println("1.Follow the sound")
		fmt.Println("2.Go to the structure")
		fmt.Scanln(&b)
		if b == 1{
			hp2 = 150
			hp2or = hp2
			attacken1 = 15
			attacken2 = 35
			hp1 = hp1or
			fmt.Println("You follow the strange noise and come upon a crystal monster")
			time.Sleep(3 * time.Second)
			fmt.Println("You enter a battle with Crystal Beast")
			for battle1 := 1 ;; battle1++ {
				if battle1 % 10 == 0 {
					fmt.Println("ultimate attack is ready")
					time.Sleep(1 * time.Second)
					fmt.Println("would you like to use ultimate attack")
					fmt.Println("1.Yes")
					fmt.Println("2.no")
					fmt.Scanln(&b)
					if b == 1 {
						time.Sleep(1 * time.Second)
						fmt.Println("using ultimate attack")
						time.Sleep(5 * time.Second)
						hp2 = hp2 - attack4
						fmt.Println("enemy took",attack4,"damage")
						if hp2 <= 0 {
							fmt.Println("You Win")
							break
						}
					} else if b == 2 {
						time.Sleep(1 * time.Second)
						continue
					} else {
						time.Sleep(1 * time.Second)
						fmt.Println("sorry that is an invalid command")
						time.Sleep(1 * time.Second)
						continue
					}
				}
				fmt.Println("your hp",hp1,"          enemy hp",hp2)
				fmt.Println("choose your attack")
				fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
				fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
				fmt.Println("3.heal",attack3,"health 100% accuracy")
				fmt.Scanln(&a)
				if a == 1 {
					time.Sleep(1 * time.Second)
					fmt.Println("enemy took",attack1,"damage")
					hp2 = hp2-attack1
					time.Sleep(1 * time.Second)
					if hp2 <= 0 {
						fmt.Println("You Win")
						break
					}
				} else if a == 2 {
					var s = time.Now().UnixNano()
					rand.Seed(s)
					var random = rand.Intn(3)
					if random == 1 || random == 2 || random == 3 {
						time.Sleep(1 * time.Second)
						hp2 = hp2 - attack2
					fmt.Println("enemy took",attack2,"damage")
					time.Sleep(1 * time.Second)
					if hp2 <= 0 {
						fmt.Println("You Win")
						break
					}
					} else {
						time.Sleep(1 * time.Second)
						fmt.Println("attack missed")
						time.Sleep(1 * time.Second)
					}
				} else if a == 3 {
					if hp1 > hp1or - attack3 {
						time.Sleep(1 * time.Second)
						fmt.Println("You're not low enough to heal")
						time.Sleep(1 * time.Second)
					} else {
						time.Sleep(1 * time.Second)
						hp1 = hp1 + attack3
						fmt.Println("You have been healed",attack3,"health")
						time.Sleep(1 * time.Second)
					}
				} else {
					fmt.Println("Sorry that is an invalid command")
				}
				var a = time.Now().UnixNano()
				rand.Seed(a)
				var random2 = rand.Intn(2)
				time.Sleep(1 * time.Second)
				fmt.Println("enemy is attacking")
				time.Sleep(1 * time.Second)
				if random2 == 0 {
					fmt.Println("you took",attacken1,"damage")
					hp1 = hp1-attacken1
					time.Sleep(1 * time.Second)
					if hp1 <= 0 {
						fmt.Println("game over")
						return
					}
				} else if random2 == 1 {
					var b = time.Now().UnixNano()
					rand.Seed(b)
					var random3 = rand.Intn(5)
						if random3 == 1 || random3 == 2 || random3 == 3 {
							fmt.Println("you took",attacken2,"damage")
							hp1 = hp1-attacken2
							time.Sleep(1 * time.Second)	
							if hp1 <= 0 {
								fmt.Println("game over")
								return
							}
						} else {
							fmt.Println("attack missed")
								time.Sleep(1 * time.Second)	
								continue
						}
					} else {
						if hp2 > hp2or - attacken3 {
							time.Sleep(1 * time.Second)
							fmt.Println("attack missed")
							time.Sleep(1 * time.Second)	
						} else {
							time.Sleep(1 * time.Second)
							hp2 = hp2 + attacken3
							fmt.Println("enemy has been healed",attacken3,"health")
							time.Sleep(1 * time.Second)
					}
				}
			}
			hp2 = hp2or
			mon = mon + 50
			time.Sleep(1 * time.Second)
			fmt.Println("You gained 50 coins")
			time.Sleep(1 * time.Second)
			fmt.Println("Total coins:",mon)
			fmt.Println("After battling the Crystal Beast you journey towards the unusual structure, and notice that it is an abandoned looking shack.")
			time.Sleep(3 * time.Second)
			fmt.Println("Confused, you take action.")
			time.Sleep(3 * time.Second)
			fmt.Println("The door appears to be locked, so you try to bust open, but the lock is too strong.")
				time.Sleep(3 * time.Second)
				fmt.Println("You notice something hanging on top of the door, nailed in by a rusty nail.")
				time.Sleep(3 * time.Second)
				fmt.Println("Head west, shall you meet the shopkeeper. Following the instructions, you head west...")
				time.Sleep(5 * time.Second)
		}else{
			fmt.Println("You journey towards the unusual structure, and notice that it is an abandoned looking shack.")
			time.Sleep(3 * time.Second)
				fmt.Println("The door appears to be locked, so you try to bust open, but the lock is too strong.")
				time.Sleep(3 * time.Second)
				fmt.Println("You notice something hanging on top of the door, nailed in by a rusty nail.")
				time.Sleep(3 * time.Second)
				fmt.Println("Head west, shall you meet the shopkeeper. Following the instructions, you head west...")
				time.Sleep(5 * time.Second)
		}
		fmt.Println("Heading west, you question your decision, whether it was wise to follow that shady note, or not.") 
		time.Sleep(3 * time.Second)
		fmt.Println("As you think, you hear leaves rustle in the distance, so you pace forward in a more hurriedly style.") 
		time.Sleep(3 * time.Second)
		fmt.Println("The sounds disappear, but you remain anxious. At long last, you stumble upon a cave, but it doesn't look normal,") 
		time.Sleep(3 * time.Second)
		fmt.Println("it actually looks familiar. You enter warily in the cave.")
		time.Sleep(3 * time.Second)
		fmt.Println("While exploring through the cave you stumble upon a Royal Gaurd")
		time.Sleep(3 * time.Second)
		fmt.Println("You enter a battle with vault gaurd")
		hp2 = 100
		hp2or = hp2
		hp1 = 100
		hp1or = hp1
		attacken1 = 15
		attacken2 = 35
		attacken3 = 20
		for battle1 := 1 ;; battle1++ {
			if battle1 % 10 == 0 {
				fmt.Println("ultimate attack is ready")
				time.Sleep(1 * time.Second)
				fmt.Println("would you like to use ultimate attack")
				fmt.Println("1.Yes")
				fmt.Println("2.no")
				fmt.Scanln(&b)
				if b == 1 {
					time.Sleep(1 * time.Second)
					fmt.Println("using ultimate attack")
					time.Sleep(5 * time.Second)
					hp2 = hp2 - attack4
					fmt.Println("enemy took",attack4,"damage")
					if hp2 <= 0 {
						fmt.Println("You Win")
						break
					}
				} else if b == 2 {
					time.Sleep(1 * time.Second)
					continue
				} else {
					time.Sleep(1 * time.Second)
					fmt.Println("sorry that is an invalid command")
					time.Sleep(1 * time.Second)
					continue
				}
			}
			fmt.Println("your hp",hp1,"          enemy hp",hp2)
			fmt.Println("choose your attack")
			fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
			fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
			fmt.Println("3.heal",attack3,"health 100% accuracy")
			fmt.Scanln(&a)
			if a == 1 {
				time.Sleep(1 * time.Second)
				fmt.Println("enemy took",attack1,"damage")
				hp2 = hp2-attack1
				time.Sleep(1 * time.Second)
				if hp2 <= 0 {
					fmt.Println("You Win")
					break
				}
			} else if a == 2 {
				var s = time.Now().UnixNano()
				rand.Seed(s)
				var random = rand.Intn(3)
				if random == 1 || random == 2 || random == 3 {
					time.Sleep(1 * time.Second)
					hp2 = hp2 - attack2
				fmt.Println("enemy took",attack2,"damage")
				time.Sleep(1 * time.Second)
				if hp2 <= 0 {
					fmt.Println("You Win")
					break
				}
				} else {
					time.Sleep(1 * time.Second)
					fmt.Println("attack missed")
					time.Sleep(1 * time.Second)
				}
			} else if a == 3 {
				if hp1 > hp1or - attack3 {
					time.Sleep(1 * time.Second)
					fmt.Println("You're not low enough to heal")
					time.Sleep(1 * time.Second)
				} else {
					time.Sleep(1 * time.Second)
					hp1 = hp1 + attack3
					fmt.Println("You have been healed",attack3,"health")
					time.Sleep(1 * time.Second)
				}
			} else {
				fmt.Println("Sorry that is an invalid command")
			}
			var a = time.Now().UnixNano()
			rand.Seed(a)
			var random2 = rand.Intn(2)
			time.Sleep(1 * time.Second)
			fmt.Println("enemy is attacking")
			time.Sleep(1 * time.Second)
			if random2 == 0 {
				fmt.Println("you took",attacken1,"damage")
				hp1 = hp1-attacken1
				time.Sleep(1 * time.Second)
				if hp1 <= 0 {
					fmt.Println("game over")
					return
				}
			} else if random2 == 1 {
				var b = time.Now().UnixNano()
				rand.Seed(b)
				var random3 = rand.Intn(5)
					if random3 == 1 || random3 == 2 || random3 == 3 {
						fmt.Println("you took",attacken2,"damage")
						hp1 = hp1-attacken2
						time.Sleep(1 * time.Second)	
						if hp1 <= 0 {
							fmt.Println("game over")
							return
						}
					} else {
						fmt.Println("attack missed")
							time.Sleep(1 * time.Second)	
							continue
					}
				} else {
					if hp2 > hp2or - attacken3 {						time.Sleep(1 * time.Second)

						fmt.Println("attack missed")
						time.Sleep(1 * time.Second)	
					} else {
						time.Sleep(1 * time.Second)
						hp2 = hp2 + attacken3
						fmt.Println("enemy has been healed",attacken3,"health")
						time.Sleep(1 * time.Second)
				}
			}
		}
		hp1 = hp1or
		mon = mon + 30
		time.Sleep(1 * time.Second)
			fmt.Println("You gained 30 coins")
			time.Sleep(1 * time.Second)
		fmt.Println("Total coins:",mon)
		time.Sleep(4 * time.Second)
		fmt.Println("After beating the Royal Guard, you notice that he was guarding a chest, without a key.")
		time.Sleep(4 * time.Second)
		fmt.Println("You open the chest and notice there is a key inside, which looks like it can open the door you saw earlier.")
		time.Sleep(4 * time.Second)
		fmt.Println("You take off east. At last, you arrive the shack, and attempt to unlock the door.") 
		time.Sleep(4 * time.Second)
		fmt.Println("The door unlocks, and you enter the shack. There is an old Elf sitting on a rocking chair, with merchandise like weapons and armor in front of him.")
		time.Sleep(4 * time.Second)
		fmt.Println("Greetings, Ashen One, he states. Confused, you greet him back. He seems to know everything about you.") 
		time.Sleep(4 * time.Second)
		fmt.Println("You gain the courage to ask him what happened to your head. He looks down and sighs, You have been cursed by Hell Lord, Hades.")
		time.Sleep(7 * time.Second)
		fmt.Println("After the initial shock, you ask the merchant why he has all this merchandise in front of him.") 
		time.Sleep(4 * time.Second)
		fmt.Println("He says he sells armor and weapons to chosen ones, who don't come very often. You ask why.") 
		time.Sleep(4 * time.Second)
		fmt.Println("He states that there is an artifact that you must fill with gems, which have been spread across four different dimensions.") 
		time.Sleep(4 * time.Second)
		fmt.Println("The first gem, is guarded by the Taurus. To beat him, you must buy my wares.")
		time.Sleep(4 * time.Second)
		for {
			fmt.Println("What would you like to buy")
			fmt.Println("Total Coins:",mon)
			fmt.Println("1.Armor Upgrade: 50 coins")
			fmt.Println("2.Weapon Upgrade: 50 coins")
			fmt.Println("3.Magic Upgrade: 50 coins")
			fmt.Println("4.Exit")
			fmt.Scanln(&b)
			if b == 1 {
				if mon < 50{
					time.Sleep(1 * time.Second)
					fmt.Println("You do not have enough to buy this item")
					time.Sleep(1 * time.Second)
				}else{
					mon = mon - 50
					hp1 = hp1 + 25
					hp1or = hp1
					time.Sleep(1 * time.Second)
					fmt.Println("Armor has been upgraded")
					time.Sleep(1 * time.Second)
				}
			}else if b == 2 {
				if mon < 50{
					time.Sleep(1 * time.Second)
					fmt.Println("You do not have enough to buy this item")
					time.Sleep(1 * time.Second)
				}else{
					mon = mon - 50
					attack1 = attack1 + 5
					attack2 = attack2 + 5
					attack4 = attack4 + 10
					time.Sleep(1 * time.Second)
					fmt.Println("Weapons has been upgraded")
					time.Sleep(1 * time.Second)
				}
			}else if b == 3{
				if mon < 50{
					time.Sleep(1 * time.Second)
					fmt.Println("You do not have enough to buy this item")
					time.Sleep(1 * time.Second)
				}else{
					mon = mon - 50
					attack3 = attack3 + 10
					time.Sleep(1 * time.Second)
					fmt.Println("Magic has been upgraded")
					time.Sleep(1 * time.Second)
				}
			}else{
				break
			}
		}
		fmt.Println("Good Luck, Ashen One! he hollers, as you exit the shack.")
		time.Sleep(3 * time.Second)
		fmt.Println("You walk through the forest, finding the wherebouts of the Taurus.")
		time.Sleep(3 * time.Second)
		fmt.Println("You then see a keep in the distance gaurded by several creatures.")
		time.Sleep(3 * time.Second)
		for taurus := 1 ;; taurus++{
			fmt.Println("Where would you like to go:")
			fmt.Println("1.Wood Dweller")
			fmt.Println("2.Ancient Log")
			fmt.Println("3.Crystal beast")
			fmt.Println("4.Taurus's throne room")
			fmt.Println("5.Shop")
			fmt.Scanln(&a)
			if a == 1 {
				if key == 1 {
					attacken1 = 15
					attacken2 = 40
					hp2 = 100
					attacken3 = 20
					hp2or = hp2
					hp1 = hp1or
					key++
					fmt.Println("Departing in 3 seconds")
					time.Sleep(3 * time.Second)
					fmt.Println("You enter a battle with Wood Dweller")
					for battle1 := 1 ;; battle1++ {
						if battle1 % 10 == 0 {
							fmt.Println("ultimate attack is ready")
							time.Sleep(1 * time.Second)
							fmt.Println("would you like to use ultimate attack")
							fmt.Println("1.Yes")
							fmt.Println("2.no")
							fmt.Scanln(&b)
							if b == 1 {
								time.Sleep(1 * time.Second)
								fmt.Println("using ultimate attack")
								time.Sleep(5 * time.Second)
								hp2 = hp2 - attack4
								fmt.Println("enemy took",attack4,"damage")
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}
							} else if b == 2 {
								time.Sleep(1 * time.Second)
								continue
							} else {
								time.Sleep(1 * time.Second)
								fmt.Println("sorry that is an invalid command")
								time.Sleep(1 * time.Second)
								continue
							}
						}
						fmt.Println("your hp",hp1,"          enemy hp",hp2)
						fmt.Println("choose your attack")
						fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
						fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
						fmt.Println("3.heal",attack3,"health 100% accuracy")
						fmt.Scanln(&a)
						if a == 1 {
							time.Sleep(1 * time.Second)
							fmt.Println("Wood Dweller took",attack1,"damage")
							hp2 = hp2-attack1
							time.Sleep(1 * time.Second)
							if hp2 <= 0 {
								fmt.Println("You Win")
								break
							}
						} else if a == 2 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(3)
							if random == 1 || random == 2 || random == 3 {
								time.Sleep(1 * time.Second)
								hp2 = hp2 - attack2
							fmt.Println("Wood Dweller took",attack2,"damage")
							time.Sleep(1 * time.Second)
							if hp2 <= 0 {
								fmt.Println("You Win")
								break
							}
							} else {
								time.Sleep(1 * time.Second)
								fmt.Println("attack missed")
								time.Sleep(1 * time.Second)
							}
						} else if a == 3 {
							if hp1 > hp1or - attack3{
								time.Sleep(1 * time.Second)
								fmt.Println("You're not low enough to heal")
								time.Sleep(1 * time.Second)
							} else {
								time.Sleep(1 * time.Second)
								hp1 = hp1 + attack3
								fmt.Println("You have been healed",attack3,"health")
								time.Sleep(1 * time.Second)
							}
						} else {
							fmt.Println("Sorry that is an invalid command")
						}
						var a = time.Now().UnixNano()
						rand.Seed(a)
						var random2 = rand.Intn(2)
						time.Sleep(1 * time.Second)
						fmt.Println("enemy is attacking")
						time.Sleep(1 * time.Second)
						if random2 == 0 {
							fmt.Println("you took",attacken1,"damage")
							hp1 = hp1-attacken1
							time.Sleep(1 * time.Second)
							if hp1 <= 0 {
								fmt.Println("game over")
								return
							}
						}else if random2 == 1 {
							var b = time.Now().UnixNano()
							rand.Seed(b)
							var random3 = rand.Intn(5)
								if random3 == 1 || random3 == 2 || random3 == 3 {
									fmt.Println("you took",attacken2,"damage")
									hp1 = hp1-attacken2
									time.Sleep(1 * time.Second)	
									if hp1 <= 0 {
										fmt.Println("game over")
										return
									}
								} else {
									fmt.Println("attack missed")
										time.Sleep(1 * time.Second)	
										continue
								}
							} else {
								if hp2 > hp2or - attacken3 {
									time.Sleep(1 * time.Second)
									fmt.Println("attack missed")
									time.Sleep(1 * time.Second)	
								} else {
									time.Sleep(1 * time.Second)
									hp2 = hp2 + attacken3
									fmt.Println("enemy has been healed",attacken3,"health")
									time.Sleep(1 * time.Second)
								}
							}
						}
						hp1 = hp1or
						time.Sleep(1 * time.Second)
						mon = mon + 50
						time.Sleep(1 * time.Second)
			fmt.Println("You gained 50 coins")
			time.Sleep(1 * time.Second)
				}else{
					fmt.Println("You already explored here. departing in 3 seconds")
					time.Sleep(3 * time.Second)
				}
			}else if a == 2 {
				if key2 == 1 {
					key2 ++
					hp1 = hp1or
					attacken1 = 20
					attacken2 = 45
					hp2 = 120
					hp2or = hp2
					attacken3 = 25
					fmt.Println("Departing in 3 seconds")
					time.Sleep(3 * time.Second)
					fmt.Println("You enter a battle with Ancient Log")
					for battle1 := 1 ;; battle1++ {
						if battle1 % 10 == 0 {
							fmt.Println("ultimate attack is ready")
							time.Sleep(1 * time.Second)
							fmt.Println("would you like to use ultimate attack")
							fmt.Println("1.Yes")
							fmt.Println("2.no")
							fmt.Scanln(&b)
							if b == 1 {
								time.Sleep(1 * time.Second)
								fmt.Println("using ultimate attack")
								time.Sleep(5 * time.Second)
								hp2 = hp2 - attack4
								fmt.Println("enemy took",attack4,"damage")
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}
							} else if b == 2 {
								time.Sleep(1 * time.Second)
								continue
							} else {
								time.Sleep(1 * time.Second)
								fmt.Println("sorry that is an invalid command")
								time.Sleep(1 * time.Second)
								continue
							}
						}
						fmt.Println("your hp",hp1,"          enemy hp",hp2)
						fmt.Println("choose your attack")
						fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
						fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
						fmt.Println("3.heal",attack3,"health 100% accuracy")
						fmt.Scanln(&a)
						if a == 1 {
							time.Sleep(1 * time.Second)
							fmt.Println("enemy took",attack1,"damage")
							hp2 = hp2-attack1
							time.Sleep(1 * time.Second)
							if hp2 <= 0 {
								fmt.Println("You Win")
								break
							}
						} else if a == 2 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(3)
							if random == 1 || random == 2 || random == 3 {
								time.Sleep(1 * time.Second)
								hp2 = hp2 - attack2
							fmt.Println("enemy took",attack2,"damage")
							time.Sleep(1 * time.Second)
							if hp2 <= 0 {
								fmt.Println("You Win")
								break
							}
							} else {
								time.Sleep(1 * time.Second)
								fmt.Println("attack missed")
								time.Sleep(1 * time.Second)
							}
						} else if a == 3 {
							if hp1 > hp1or - attack3 {
								time.Sleep(1 * time.Second)
								fmt.Println("You're not low enough to heal")
								time.Sleep(1 * time.Second)
							} else {
								time.Sleep(1 * time.Second)
								hp1 = hp1 + attack3
								fmt.Println("You have been healed",attack3,"health")
								time.Sleep(1 * time.Second)
							}
						} else {
							fmt.Println("Sorry that is an invalid command")
						}
						var a = time.Now().UnixNano()
						rand.Seed(a)
						var random2 = rand.Intn(2)
						time.Sleep(1 * time.Second)
						fmt.Println("enemy is attacking")
						time.Sleep(1 * time.Second)
						if random2 == 0 {
							fmt.Println("you took",attacken1,"damage")
							hp1 = hp1-attacken1
							time.Sleep(1 * time.Second)
							if hp1 <= 0 {
								fmt.Println("game over")
								return
							}
						} else if random2 == 1 {
							var b = time.Now().UnixNano()
							rand.Seed(b)
							var random3 = rand.Intn(5)
								if random3 == 1 || random3 == 2 || random3 == 3 {
									fmt.Println("you took",attacken2,"damage")
									hp1 = hp1-attacken2
									time.Sleep(1 * time.Second)	
									if hp1 <= 0 {
										fmt.Println("game over")
										return
									}
								} else {
									fmt.Println("attack missed")
										time.Sleep(1 * time.Second)	
										continue
								}
							} else {
								if hp2 > hp2or - attacken3 {
									time.Sleep(1 * time.Second)
									fmt.Println("attack missed")
									time.Sleep(1 * time.Second)	
								} else {
									time.Sleep(1 * time.Second)
									hp2 = hp2 + attacken3
									fmt.Println("enemy has been healed",attacken3,"health")
									time.Sleep(1 * time.Second)
							}
						}
					}
					hp1 = hp1or
					hp1 = hp1or
					mon = mon + 50
					fmt.Println("You gained 50 coins")
					time.Sleep(1 * time.Second)
					fmt.Println("Total coins:",mon)
					time.Sleep(1 * time.Second)
				}else{
					fmt.Println("You already explored here. departing in 3 seconds")
					time.Sleep(3 * time.Second)
				}
			}else if a == 3 {
				if key3 == 1 {
					hp2 = hp2or
					key3 ++
					attacken1 = 25
					attacken2 = 45
					hp2 = 160
					attacken3 = 30
					hp1 = hp1or
					fmt.Println("Departing in 3 seconds")
					time.Sleep(3 * time.Second)
					fmt.Println("You enter a battle with Crystal Beast")
					for battle1 := 1 ;; battle1++ {
						if battle1 % 10 == 0 {
							fmt.Println("ultimate attack is ready")
							time.Sleep(1 * time.Second)
							fmt.Println("would you like to use ultimate attack")
							fmt.Println("1.Yes")
							fmt.Println("2.no")
							fmt.Scanln(&b)
							if b == 1 {
								time.Sleep(1 * time.Second)
								fmt.Println("using ultimate attack")
								time.Sleep(5 * time.Second)
								hp2 = hp2 - attack4
								fmt.Println("enemy took",attack4,"damage")
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}
							} else if b == 2 {
								time.Sleep(1 * time.Second)
								continue
							} else {
								time.Sleep(1 * time.Second)
								fmt.Println("sorry that is an invalid command")
								time.Sleep(1 * time.Second)
								continue
							}
						}
						fmt.Println("your hp",hp1,"          enemy hp",hp2)
						fmt.Println("choose your attack")
						fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
						fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
						fmt.Println("3.heal",attack3,"health 100% accuracy")
						fmt.Scanln(&a)
						if a == 1 {
							time.Sleep(1 * time.Second)
							fmt.Println("enemy took",attack1,"damage")
							hp2 = hp2-attack1
							time.Sleep(1 * time.Second)
							if hp2 <= 0 {
								fmt.Println("You Win")
								break
							}
						} else if a == 2 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(3)
							if random == 1 || random == 2 || random == 3 {
								time.Sleep(1 * time.Second)
								hp2 = hp2 - attack2
							fmt.Println("enemy took",attack2,"damage")
							time.Sleep(1 * time.Second)
							if hp2 <= 0 {
								fmt.Println("You Win")
								break
							}
							} else {
								time.Sleep(1 * time.Second)
								fmt.Println("attack missed")
								time.Sleep(1 * time.Second)
							}
						} else if a == 3 {
							if hp1 > hp1or - attack3 {
								time.Sleep(1 * time.Second)
								fmt.Println("You're not low enough to heal")
								time.Sleep(1 * time.Second)
							} else {
								time.Sleep(1 * time.Second)
								hp1 = hp1 + attack3
								fmt.Println("You have been healed",attack3,"health")
								time.Sleep(1 * time.Second)
							}
						} else {
							fmt.Println("Sorry that is an invalid command")
						}
						var a = time.Now().UnixNano()
						rand.Seed(a)
						var random2 = rand.Intn(2)
						time.Sleep(1 * time.Second)
						fmt.Println("enemy is attacking")
						time.Sleep(1 * time.Second)
						if random2 == 0 {
							fmt.Println("you took",attacken1,"damage")
							hp1 = hp1-attacken1
							time.Sleep(1 * time.Second)
							if hp1 <= 0 {
								fmt.Println("game over")
								return
							}
						} else if random2 == 1 {
							var b = time.Now().UnixNano()
							rand.Seed(b)
							var random3 = rand.Intn(5)
								if random3 == 1 || random3 == 2 || random3 == 3 {
									fmt.Println("you took",attacken2,"damage")
									hp1 = hp1-attacken2
									time.Sleep(1 * time.Second)	
									if hp1 <= 0 {
										fmt.Println("game over")
										return
									}
								} else {
									fmt.Println("attack missed")
										time.Sleep(1 * time.Second)	
										continue
								}
							} else {
								if hp2 > hp2or - attacken3 {
									time.Sleep(1 * time.Second)
									fmt.Println("attack missed")
									time.Sleep(1 * time.Second)	
								} else {
									time.Sleep(1 * time.Second)
									hp2 = hp2 + attacken3
									fmt.Println("enemy has been healed",attacken3,"health")
									time.Sleep(1 * time.Second)
							}
						}
					}
					hp1 = hp1or
					mon = mon + 50
					time.Sleep(1 * time.Second)
				fmt.Println("You gained 50 coins")
				time.Sleep(1 * time.Second)
				fmt.Println("The Beast drops to the ground grasping a note in his hand: 201514")
				}else{
					fmt.Println("You already explored here. departing in 3 seconds")
					time.Sleep(3 * time.Second)
				}
			}else if a == 4 {
				time.Sleep(1 * time.Second)
				fmt.Println("Please enter the code")
				fmt.Scanln(&b)
				if b == 201514 {
					fmt.Println("You enter a dungeon with a large monster weilding a giant axe")
					time.Sleep(1 * time.Second)
					fmt.Println("You enter a battle with Taurus")
					hp2 = hp2or
					key3 ++
					attacken1 = 30
					attacken2 = 60
					hp2 = 220
					attacken3 = 30
					hp1 = hp1or
					for battle1 := 1 ;; battle1++ {
						if battle1 % 10 == 0 {
							fmt.Println("ultimate attack is ready")
							time.Sleep(1 * time.Second)
							fmt.Println("would you like to use ultimate attack")
							fmt.Println("1.Yes")
							fmt.Println("2.no")
							fmt.Scanln(&b)
							if b == 1 {
								time.Sleep(1 * time.Second)
								fmt.Println("using ultimate attack")
								time.Sleep(5 * time.Second)
								hp2 = hp2 - attack4
								fmt.Println("enemy took",attack4,"damage")
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}
							} else if b == 2 {
								time.Sleep(1 * time.Second)
								continue
							} else {
								time.Sleep(1 * time.Second)
								fmt.Println("sorry that is an invalid command")
								time.Sleep(1 * time.Second)
								continue
							}
						}
						fmt.Println("your hp",hp1,"          enemy hp",hp2)
						fmt.Println("choose your attack")
						fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
						fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
						fmt.Println("3.heal",attack3,"health 100% accuracy")
						fmt.Scanln(&a)
						if a == 1 {
							time.Sleep(1 * time.Second)
							fmt.Println("enemy took",attack1,"damage")
							hp2 = hp2-attack1
							time.Sleep(1 * time.Second)
							if hp2 <= 0 {
								fmt.Println("You Win")
								break
							}
						} else if a == 2 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(3)
							if random == 1 || random == 2 || random == 3 {
								time.Sleep(1 * time.Second)
								hp2 = hp2 - attack2
							fmt.Println("enemy took",attack2,"damage")
							time.Sleep(1 * time.Second)
							if hp2 <= 0 {
								fmt.Println("You Win")
								break
							}
							} else {
								time.Sleep(1 * time.Second)
								fmt.Println("attack missed")
								time.Sleep(1 * time.Second)
							}
						} else if a == 3 {
							if hp1 > hp1or - attack3 {
								time.Sleep(1 * time.Second)
								fmt.Println("You're not low enough to heal")
								time.Sleep(1 * time.Second)
							} else {
								time.Sleep(1 * time.Second)
								hp1 = hp1 + attack3
								fmt.Println("You have been healed",attack3,"health")
								time.Sleep(1 * time.Second)
							}
						} else {
							fmt.Println("Sorry that is an invalid command")
						}
						var a = time.Now().UnixNano()
						rand.Seed(a)
						var random2 = rand.Intn(2)
						time.Sleep(1 * time.Second)
						fmt.Println("enemy is attacking")
						time.Sleep(1 * time.Second)
						if random2 == 0 {
							fmt.Println("you took",attacken1,"damage")
							hp1 = hp1-attacken1
							time.Sleep(1 * time.Second)
							if hp1 <= 0 {
								fmt.Println("game over")
								return
							}
						} else if random2 == 1 {
							var b = time.Now().UnixNano()
							rand.Seed(b)
							var random3 = rand.Intn(5)
								if random3 == 1 || random3 == 2 || random3 == 3 {
									fmt.Println("you took",attacken2,"damage")
									hp1 = hp1-attacken2
									time.Sleep(1 * time.Second)	
									if hp1 <= 0 {
										fmt.Println("game over")
										return
									}
								} else {
									fmt.Println("attack missed")
										time.Sleep(1 * time.Second)	
										continue
								}
							} else {
								if hp2 > hp2or - attacken3 {
									time.Sleep(1 * time.Second)
									fmt.Println("attack missed")
									time.Sleep(1 * time.Second)	
								} else {
									time.Sleep(1 * time.Second)
									hp2 = hp2 + attacken3
									fmt.Println("enemy has been healed",attacken3,"health")
									time.Sleep(1 * time.Second)
							}
						}
					}
					hp1 = hp1or
					time.Sleep(4 * time.Second)
					mon = mon + 200
					time.Sleep(1 * time.Second)
				fmt.Println("You gained 200 coins")
				time.Sleep(1 * time.Second)
					fmt.Println("After beating the Taurus, you scavenge its carcass, and find a key, which opens the chest.")
					time.Sleep(4 * time.Second)
					fmt.Println("The first, bright yellow gem is inside. You have saved Mystic Woods from the curse.")
					time.Sleep(4 * time.Second)
					fmt.Println("Now, you must go, and conquer Mount Olympus")
				time.Sleep(7 * time.Second)
				break
				}else{
					time.Sleep(1 * time.Second)
					fmt.Println("That is incorrect")
					time.Sleep(1 * time.Second)
				}
			}else{
				time.Sleep(1 * time.Second)
				for {
					fmt.Println("What would you like to buy")
					fmt.Println("Total Coins:",mon)
					fmt.Println("1.Armor Upgrade: 50 coins")
					fmt.Println("2.Weapon Upgrade: 50 coins")
					fmt.Println("3.Magic Upgrade: 50 coins")
					fmt.Println("4.Exit")
					fmt.Scanln(&b)
					if b == 1 {
						if mon < 50{
							time.Sleep(1 * time.Second)
							fmt.Println("You do not have enough to buy this item")
							time.Sleep(1 * time.Second)
						}else{
							mon = mon - 50
							hp1 = hp1 + 25
							hp1or = hp1
							time.Sleep(1 * time.Second)
							fmt.Println("Armor has been upgraded")
							time.Sleep(1 * time.Second)
						}
					}else if b == 2 {
						if mon < 50{
							time.Sleep(1 * time.Second)
							fmt.Println("You do not have enough to buy this item")
							time.Sleep(1 * time.Second)
						}else{
							mon = mon - 50
							attack1 = attack1 + 10
							attack2 = attack2 + 10
							attack4 = attack4 + 15
							time.Sleep(1 * time.Second)
							fmt.Println("Weapons has been upgraded")
							time.Sleep(1 * time.Second)
						}
					}else if b == 3{
						if mon < 50{
							time.Sleep(1 * time.Second)
							fmt.Println("You do not have enough to buy this item")
							time.Sleep(1 * time.Second)
						}else{
							mon = mon - 50
							attack3 = attack3 + 10
							time.Sleep(1 * time.Second)
							fmt.Println("Magic has been upgraded")
							time.Sleep(1 * time.Second)
						}
					}else{
						time.Sleep(1 * time.Second)
						break
					}
				}
			}
		}
		for space := 1 ; space <=100 ; space++{
			fmt.Println()
		}
		fmt.Println("MAP 2: Mount Olympus")
		for space := 1 ; space <=20 ; space++{
			fmt.Println()
		}
		time.Sleep(2 * time.Second)	
		for space := 1 ; space <=100 ; space++{
			fmt.Println()
		}
		time.Sleep(2 * time.Second)
		fmt.Println("The ship finally arrives at a large island, with 2 smaller mountains, and one mountain, bigger than you've ever seen.")
		time.Sleep(4 * time.Second)
		fmt.Println("The shipman puts the boat on the dock, and you climb off the wooden stairs on the platform. Travel Safe, Ashen One. The Shipman says.")
		time.Sleep(4 * time.Second)
		fmt.Println("You set off for your mission to find the Red Gem. A Lesser Jerboa scurries past as you begin your trek up the mountain. He was panicking.")
		time.Sleep(4 * time.Second)
		fmt.Println("You get ready for a battle. It is a Possessed Spartan.")
		time.Sleep(4 * time.Second)
		fmt.Println("You enter a battle with Possessed Spartan")
		hp2 = hp2or
		attacken1 = 20
		attacken2 = 40
		hp2 = 130
		attacken3 = 30
		hp1 = hp1or
		for battle1 := 1 ;; battle1++ {
			if battle1 % 10 == 0 {
				fmt.Println("ultimate attack is ready")
				time.Sleep(1 * time.Second)
				fmt.Println("would you like to use ultimate attack")
				fmt.Println("1.Yes")
				fmt.Println("2.no")
				fmt.Scanln(&b)
				if b == 1 {
					time.Sleep(1 * time.Second)
					fmt.Println("using ultimate attack")
					time.Sleep(5 * time.Second)
					hp2 = hp2 - attack4
					fmt.Println("enemy took",attack4,"damage")
					if hp2 <= 0 {
						fmt.Println("You Win")
						break
					}
				} else if b == 2 {
					time.Sleep(1 * time.Second)
					continue
				} else {
					time.Sleep(1 * time.Second)
					fmt.Println("sorry that is an invalid command")
					time.Sleep(1 * time.Second)
					continue
				}
			}
			fmt.Println("your hp",hp1,"          enemy hp",hp2)
			fmt.Println("choose your attack")
			fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
			fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
			fmt.Println("3.heal",attack3,"health 100% accuracy")
			fmt.Scanln(&a)
			if a == 1 {
				time.Sleep(1 * time.Second)
				fmt.Println("enemy took",attack1,"damage")
				hp2 = hp2-attack1
				time.Sleep(1 * time.Second)
				if hp2 <= 0 {
					fmt.Println("You Win")
					break
				}
			} else if a == 2 {
				var s = time.Now().UnixNano()
				rand.Seed(s)
				var random = rand.Intn(3)
				if random == 1 || random == 2 || random == 3 {
					time.Sleep(1 * time.Second)
					hp2 = hp2 - attack2
				fmt.Println("enemy took",attack2,"damage")
				time.Sleep(1 * time.Second)
				if hp2 <= 0 {
					fmt.Println("You Win")
					break
				}
				} else {
					time.Sleep(1 * time.Second)
					fmt.Println("attack missed")
					time.Sleep(1 * time.Second)
				}
			} else if a == 3 {
				if hp1 > hp1or - attack3 {
					time.Sleep(1 * time.Second)
					fmt.Println("You're not low enough to heal")
					time.Sleep(1 * time.Second)
				} else {
					time.Sleep(1 * time.Second)
					hp1 = hp1 + attack3
					fmt.Println("You have been healed",attack3,"health")
					time.Sleep(1 * time.Second)
				}
			} else {
				fmt.Println("Sorry that is an invalid command")
			}
			var a = time.Now().UnixNano()
			rand.Seed(a)
			var random2 = rand.Intn(2)
			time.Sleep(1 * time.Second)
			fmt.Println("enemy is attacking")
			time.Sleep(1 * time.Second)
			if random2 == 0 {
				fmt.Println("you took",attacken1,"damage")
				hp1 = hp1-attacken1
				time.Sleep(1 * time.Second)
				if hp1 <= 0 {
					fmt.Println("game over")
					return
				}
			} else if random2 == 1 {
				var b = time.Now().UnixNano()
				rand.Seed(b)
				var random3 = rand.Intn(5)
					if random3 == 1 || random3 == 2 || random3 == 3 {
						fmt.Println("you took",attacken2,"damage")
						hp1 = hp1-attacken2
						time.Sleep(1 * time.Second)	
						if hp1 <= 0 {
							fmt.Println("game over")
							return
						}
					} else {
						fmt.Println("attack missed")
							time.Sleep(1 * time.Second)	
							continue
					}
				} else {
					if hp2 > hp2or - attacken3 {
						time.Sleep(1 * time.Second)
						fmt.Println("attack missed")
						time.Sleep(1 * time.Second)	
					} else {
						time.Sleep(1 * time.Second)
						hp2 = hp2 + attacken3
						fmt.Println("enemy has been healed",attacken3,"health")
						time.Sleep(1 * time.Second)
				}
			}
		}
		hp1 = hp1or
		hp1 = hp1 + 10
		time.Sleep(1 * time.Second)
		fmt.Println("After the battle, you claim its helmet.")
		time.Sleep(4 * time.Second)
		fmt.Println("You know what's coming for you. There is no good left in this Mountain range. You are aware that you must conquer two other mountains before Olympus.")
		time.Sleep(4 * time.Second)
		fmt.Println("The first mountain, is merely  just a hill. On the top of the mountain, waits the Brood Mother, a lethally dangerous spider.")
		time.Sleep(4 * time.Second)
		fmt.Println("When you hike up the hill, you notice some movement. It is a possessed spartan.  As soon as you notice it, the spartan gets impaled by a ten-foot spider.")
		time.Sleep(4 * time.Second)
		fmt.Println("You run away into a bush, hiding.")
	}