package main

import (
	"fmt"
	"math/rand"
	"time"
)
func main() {
	var game int
	for {
		fmt.Println("Welcome to Demons Gate. What mode would you like to play:")
		fmt.Println("1.Campaign")
		fmt.Println("2.Playground")
		fmt.Println("3.PvP")
		fmt.Println("4.Survival")
		fmt.Println("5.Campaign 2")
		fmt.Println("6.Instructions")
		fmt.Scanln(&game)
		fmt.Println("Starting")
		time.Sleep(5 * time.Second)
		if game ==1 {
			var a int
			var check = 1
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
			var hp1tr int
			var attack1tr int
			var attack2tr int
			var attack3tr int
			var attack4tr int
			var attacken1 = 10
			var attacken2 = 20 
			var attacken3 = 10
			var montr int
			var check4 int
			for {
				attack1 = attack1 - attack1tr
				mon = mon - montr
				hp1 = hp1 - hp1tr
				attack2 = attack2 - attack2tr
				attack3 = attack3 - attack3tr
				attack4 = attack4 - attack4tr
				hp1 = hp1or
				hp2 = 75
				attacken1 = 10
				attacken2 = 20
				attacken3 = 10
				hp2or = hp2
				attack1tr = 0
				hp1tr = 0
				attack2tr = 0
				attack3tr = 0
				attack4tr = 0
				montr = 0
				check = 1
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
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(8)
							if random == 3 {
								hp2 = hp2 - attack4 * 3 / 2
								fmt.Println("Critical Hit")
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}	
							}else {
								hp2 = hp2 - attack4
								fmt.Println("Enemy took",attack4,"damage")
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}
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
					fmt.Println()
					fmt.Println("choose your attack")
					fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
					fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
					fmt.Println("3.heal",attack3,"health 100% accuracy")
					fmt.Scanln(&a)
					if a == 1 {
						var s = time.Now().UnixNano()
						rand.Seed(s)
						var random = rand.Intn(8)
						if random == 0 {
							hp2 = hp2 - attack1 * 3 / 2
							fmt.Println("Critical Hit")
							time.Sleep(1 * time.Second)
							fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
							time.Sleep(1 * time.Second)
							if hp2 <= 0 {
								fmt.Println("You Win")
								break
							}	
						}else {
							time.Sleep(1 * time.Second)
							fmt.Println("Enemy took",attack1,"damage")
							hp2 = hp2-attack1
							time.Sleep(1 * time.Second)
							if hp2 <= 0 {
								fmt.Println("You Win")
								break
							}
						}
					} else if a == 2 {
						var s = time.Now().UnixNano()
						rand.Seed(s)
						var random = rand.Intn(3)
						if random == 1 || random == 2 || random == 3 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(8)
							if random == 3 {
								hp2 = hp2 - attack2 * 3 / 2
								fmt.Println("Critical Hit")
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}	
							}else {
							time.Sleep(1 * time.Second)
							hp2 = hp2 - attack2
							fmt.Println("Enemy took",attack2,"damage")
							time.Sleep(1 * time.Second)
							if hp2 <= 0 {
								fmt.Println("You Win")
								break
							}
							}
						} else {
							time.Sleep(1 * time.Second)
							fmt.Println("attack missed")
							time.Sleep(1 * time.Second)
						}
					} else if a == 3 {
						if hp1 == hp1or {
							time.Sleep(1 * time.Second)
							fmt.Println("You're not low enough to heal")
							time.Sleep(1 * time.Second)
						} else if hp1 >= (hp1or - attack3){
							time.Sleep(1 * time.Second)
							hp1 = hp1or
							fmt.Println("You have been healed to max")
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
					time.Sleep(1 * time.Second)
					fmt.Println("enemy is attacking")
					time.Sleep(1 * time.Second)
					if hp2 <= (hp2or / 2) && hp1 > attacken2 {
						time.Sleep(1 * time.Second)
						hp2 = hp2 + attacken3
						fmt.Println("enemy has been healed",attacken3,"health")
						time.Sleep(1 * time.Second)
						continue
					}
					if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
						var s = time.Now().UnixNano()
						rand.Seed(s)
						var random = rand.Intn(8)
						if random == 0 {
							hp1 = hp1 - attacken1 * 3 / 2
							fmt.Println("Critical Hit")
							time.Sleep(1 * time.Second)
							fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
							time.Sleep(1 * time.Second)
							if hp1 <= 0 {
								fmt.Println("You Died")
								check = 2
								break
							}	
						}else {
							fmt.Println("you took",attacken1,"damage")
							hp1 = hp1-attacken1
							time.Sleep(1 * time.Second)
							if hp1 <= 0 {
								fmt.Println("You Died")
								check = 2
								break
							}
						}
					} else {
						var b = time.Now().UnixNano()
						rand.Seed(b)
						var random3 = rand.Intn(5)
							if random3 == 1 || random3 == 2 || random3 == 3 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 0 {
									hp1 = hp1 - attacken2 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
									time.Sleep(1 * time.Second)
									if hp1 <= 0 {
										fmt.Println("You Died")
										check = 2
										break
									}	
								}else {
									fmt.Println("you took",attacken2,"damage")
									hp1 = hp1-attacken2
									time.Sleep(1 * time.Second)	
									if hp1 <= 0 {
										fmt.Println("You Died")
										check = 2
										break
									}
								}
							} else {
								fmt.Println("attack missed")
									time.Sleep(1 * time.Second)	
									continue
							}
						}
					}
					if check == 2 {
						time.Sleep(2 * time.Second)
						fmt.Println("Respawning")
						time.Sleep(5 * time.Second)
						continue
					}
					time.Sleep(2 * time.Second)
					fmt.Println("The Wood Dweller drops dead, and you are victorious. You scavenge its body and claim its weapon")
					time.Sleep(2 * time.Second)
					attack1 = attack1 + 10
					attack2 = attack2 + 10
					attack4 = attack4 + 10
					attack1tr = attack1tr + 10
					attack2tr = attack2tr + 10
					attack4tr = attack4tr + 10
					fmt.Println("Weapon upgraded")
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
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 3 {
										hp2 = hp2 - attack4 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									}else {
										hp2 = hp2 - attack4
										fmt.Println("Enemy took",attack4,"damage")
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
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
							fmt.Println()
							fmt.Println("choose your attack")
							fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
							fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
							fmt.Println("3.heal",attack3,"health 100% accuracy")
							fmt.Scanln(&a)
							if a == 1 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 0 {
									hp2 = hp2 - attack1 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								} else {
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack1,"damage")
									hp2 = hp2-attack1
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}
								}
							} else if a == 2 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(3)
								if random == 1 || random == 2 || random == 3 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 3 {
										hp2 = hp2 - attack2 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									} else {
									time.Sleep(1 * time.Second)
									hp2 = hp2 - attack2
									fmt.Println("Enemy took",attack2,"damage")
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}
									}
								} else {
									time.Sleep(1 * time.Second)
									fmt.Println("attack missed")
									time.Sleep(1 * time.Second)
								}
							} else if a == 3 {
								if hp1 == hp1or {
									time.Sleep(1 * time.Second)
									fmt.Println("You're not low enough to heal")
									time.Sleep(1 * time.Second)
								} else if hp1 >= (hp1or - attack3){
									time.Sleep(1 * time.Second)
									hp1 = hp1or
									fmt.Println("You have been healed to max")
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
							time.Sleep(1 * time.Second)
							fmt.Println("enemy is attacking")
							time.Sleep(1 * time.Second)
							if hp2 <= (hp2or / 2) && hp1 > attacken2 {
								time.Sleep(1 * time.Second)
								hp2 = hp2 + attacken3
								fmt.Println("enemy has been healed",attacken3,"health")
								time.Sleep(1 * time.Second)
								continue
							}
							if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 0 {
									hp1 = hp1 - attacken1 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
									time.Sleep(1 * time.Second)
									if hp1 <= 0 {
										fmt.Println("You Died")
										check = 2
										break
									}	
								}else {
									fmt.Println("you took",attacken1,"damage")
									hp1 = hp1-attacken1
									time.Sleep(1 * time.Second)
									if hp1 <= 0 {
										fmt.Println("You Died")
										check = 2
										break
									}
								}
							} else {
								var b = time.Now().UnixNano()
								rand.Seed(b)
								var random3 = rand.Intn(5)
									if random3 == 1 || random3 == 2 || random3 == 3 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp1 = hp1 - attacken2 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}	
										}else {
											fmt.Println("you took",attacken2,"damage")
											hp1 = hp1-attacken2
											time.Sleep(1 * time.Second)	
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}
										}
									} else {
										fmt.Println("attack missed")
											time.Sleep(1 * time.Second)	
											continue
									}
								}
							}
						if check == 2 {
							time.Sleep(2 * time.Second)
							fmt.Println("Respawning")
							time.Sleep(5 * time.Second)
							continue
						}
						attack1tr = 0
						hp1tr = 0
						attack2tr = 0
						attack3tr = 0
						attack4tr = 0
						montr = 0
						break
			}
				time.Sleep(1 * time.Second)
				fmt.Println("You gained 20 coins")
				time.Sleep(1 * time.Second)
				mon = mon + 20
				fmt.Println("Total coins:",mon)
				time.Sleep(3 * time.Second)
			for {
				attack1 = attack1 - attack1tr
				mon = mon - montr
				hp1 = hp1 - hp1tr
				attack2 = attack2 - attack2tr
				attack3 = attack3 - attack3tr
				attack4 = attack4 - attack4tr
				hp1 = hp1or
				hp2 = hp2or
				attack1tr = 0
				hp1tr = 0
				attack2tr = 0
				attack3tr = 0
				attack4tr = 0
				montr = 0
				check = 1
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
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 3 {
									hp2 = hp2 - attack4 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}else {
									hp2 = hp2 - attack4
									fmt.Println("Enemy took",attack4,"damage")
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}
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
						fmt.Println()
						fmt.Println("choose your attack")
						fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
						fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
						fmt.Println("3.heal",attack3,"health 100% accuracy")
						fmt.Scanln(&a)
						if a == 1 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(8)
							if random == 0 {
								hp2 = hp2 - attack1 * 3 / 2
								fmt.Println("Critical Hit")
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}	
							}else {
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy took",attack1,"damage")
								hp2 = hp2-attack1
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}
							}
						} else if a == 2 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(3)
							if random == 1 || random == 2 || random == 3 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 3 {
									hp2 = hp2 - attack2 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}else {
								time.Sleep(1 * time.Second)
								hp2 = hp2 - attack2
								fmt.Println("Enemy took",attack2,"damage")
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}
								}
							} else {
								time.Sleep(1 * time.Second)
								fmt.Println("attack missed")
								time.Sleep(1 * time.Second)
							}
						} else if a == 3 {
							if hp1 == hp1or {
								time.Sleep(1 * time.Second)
								fmt.Println("You're not low enough to heal")
								time.Sleep(1 * time.Second)
							} else if hp1 >= (hp1or - attack3){
								time.Sleep(1 * time.Second)
								hp1 = hp1or
								fmt.Println("You have been healed to max")
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
						time.Sleep(1 * time.Second)
						fmt.Println("enemy is attacking")
						time.Sleep(1 * time.Second)
						if hp2 <= (hp2or / 2) && hp1 > attacken2 {
							time.Sleep(1 * time.Second)
							hp2 = hp2 + attacken3
							fmt.Println("enemy has been healed",attacken3,"health")
							time.Sleep(1 * time.Second)
							continue
						}
						if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(8)
							if random == 0 {
								hp1 = hp1 - attacken1 * 3 / 2
								fmt.Println("Critical Hit")
								time.Sleep(1 * time.Second)
								fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
								time.Sleep(1 * time.Second)
								if hp1 <= 0 {
									fmt.Println("You Died")
									check = 2
									break
								}	
							}else {
								fmt.Println("you took",attacken1,"damage")
								hp1 = hp1-attacken1
								time.Sleep(1 * time.Second)
								if hp1 <= 0 {
									fmt.Println("You Died")
									check = 2
									break
								}
							}
						} else {
							var b = time.Now().UnixNano()
							rand.Seed(b)
							var random3 = rand.Intn(5)
								if random3 == 1 || random3 == 2 || random3 == 3 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp1 = hp1 - attacken2 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											check = 2
											break
										}	
									}else {
										fmt.Println("you took",attacken2,"damage")
										hp1 = hp1-attacken2
										time.Sleep(1 * time.Second)	
										if hp1 <= 0 {
											fmt.Println("You Died")
											check = 2
											break
										}
									}
								} else {
									fmt.Println("attack missed")
										time.Sleep(1 * time.Second)	
										continue
								}
							}
						}
					if check == 2 {
						time.Sleep(2 * time.Second)
						fmt.Println("Respawning")
						time.Sleep(5 * time.Second)
						continue
					}
					hp2 = hp2or
					mon = mon + 50
					montr = montr + 50
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
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(8)
							if random == 3 {
								hp2 = hp2 - attack4 * 3 / 2
								fmt.Println("Critical Hit")
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}	
							}else {
								hp2 = hp2 - attack4
								fmt.Println("Enemy took",attack4,"damage")
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}
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
					fmt.Println()
					fmt.Println("choose your attack")
					fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
					fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
					fmt.Println("3.heal",attack3,"health 100% accuracy")
					fmt.Scanln(&a)
					if a == 1 {
						var s = time.Now().UnixNano()
						rand.Seed(s)
						var random = rand.Intn(8)
						if random == 0 {
							hp2 = hp2 - attack1 * 3 / 2
							fmt.Println("Critical Hit")
							time.Sleep(1 * time.Second)
							fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
							time.Sleep(1 * time.Second)
							if hp2 <= 0 {
								fmt.Println("You Win")
								break
							}	
						}else {
							time.Sleep(1 * time.Second)
							fmt.Println("Enemy took",attack1,"damage")
							hp2 = hp2-attack1
							time.Sleep(1 * time.Second)
							if hp2 <= 0 {
								fmt.Println("You Win")
								break
							}
						}
					} else if a == 2 {
						var s = time.Now().UnixNano()
						rand.Seed(s)
						var random = rand.Intn(3)
						if random == 1 || random == 2 || random == 3 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(8)
							if random == 3 {
								hp2 = hp2 - attack2 * 3 / 2
								fmt.Println("Critical Hit")
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}	
							}else {
							time.Sleep(1 * time.Second)
							hp2 = hp2 - attack2
							fmt.Println("Enemy took",attack2,"damage")
							time.Sleep(1 * time.Second)
							if hp2 <= 0 {
								fmt.Println("You Win")
								break
							}
							}
						} else {
							time.Sleep(1 * time.Second)
							fmt.Println("attack missed")
							time.Sleep(1 * time.Second)
						}
					} else if a == 3 {
						if hp1 == hp1or {
							time.Sleep(1 * time.Second)
							fmt.Println("You're not low enough to heal")
							time.Sleep(1 * time.Second)
						} else if hp1 >= (hp1or - attack3){
							time.Sleep(1 * time.Second)
							hp1 = hp1or
							fmt.Println("You have been healed to max")
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
					time.Sleep(1 * time.Second)
					fmt.Println("enemy is attacking")
					time.Sleep(1 * time.Second)
					if hp2 <= (hp2or / 2) && hp1 > attacken2 {
						time.Sleep(1 * time.Second)
						hp2 = hp2 + attacken3
						fmt.Println("enemy has been healed",attacken3,"health")
						time.Sleep(1 * time.Second)
						continue
					}
					if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
						var s = time.Now().UnixNano()
						rand.Seed(s)
						var random = rand.Intn(8)
						if random == 0 {
							hp1 = hp1 - attacken1 * 3 / 2
							fmt.Println("Critical Hit")
							time.Sleep(1 * time.Second)
							fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
							time.Sleep(1 * time.Second)
							if hp1 <= 0 {
								fmt.Println("You Died")
								check = 2
								break
							}	
						}else {
							fmt.Println("you took",attacken1,"damage")
							hp1 = hp1-attacken1
							time.Sleep(1 * time.Second)
							if hp1 <= 0 {
								fmt.Println("You Died")
								check = 2
								break
							}
						}
					} else {
						var b = time.Now().UnixNano()
						rand.Seed(b)
						var random3 = rand.Intn(5)
							if random3 == 1 || random3 == 2 || random3 == 3 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 0 {
									hp1 = hp1 - attacken2 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
									time.Sleep(1 * time.Second)
									if hp1 <= 0 {
										fmt.Println("You Died")
										check = 2
										break
									}	
								}else {
									fmt.Println("you took",attacken2,"damage")
									hp1 = hp1-attacken2
									time.Sleep(1 * time.Second)	
									if hp1 <= 0 {
										fmt.Println("You Died")
										check = 2
										break
									}
								}
							} else {
								fmt.Println("attack missed")
									time.Sleep(1 * time.Second)	
									continue
							}
						}
					}
				if check == 2 {
					time.Sleep(2 * time.Second)
					fmt.Println("Respawning")
					time.Sleep(5 * time.Second)
					continue
				}
				attack1tr = 0
				hp1tr = 0
				attack2tr = 0
				attack3tr = 0
				attack4tr = 0
				montr = 0
				hp1 = hp1or
				mon = mon + 30
				time.Sleep(1 * time.Second)
					fmt.Println("You gained 30 coins")
					time.Sleep(1 * time.Second)
				fmt.Println("Total coins:",mon)
				time.Sleep(4 * time.Second)
				break
			}
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
					hp1 = hp1or
					hp2 = hp2or
					check = 1
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
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack4 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
												hp2 = hp2 - attack4
												fmt.Println("Enemy took",attack4,"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
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
									fmt.Println()
									fmt.Println("choose your attack")
									fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
									fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
									fmt.Println("3.heal",attack3,"health 100% accuracy")
									fmt.Scanln(&a)
									if a == 1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp2 = hp2 - attack1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1,"damage")
											hp2 = hp2-attack1
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
										}
									} else if a == 2 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(3)
										if random == 1 || random == 2 || random == 3 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack2 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
											time.Sleep(1 * time.Second)
											hp2 = hp2 - attack2
											fmt.Println("Enemy took",attack2,"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
											}
										} else {
											time.Sleep(1 * time.Second)
											fmt.Println("attack missed")
											time.Sleep(1 * time.Second)
										}
									} else if a == 3 {
										if hp1 == hp1or {
											time.Sleep(1 * time.Second)
											fmt.Println("You're not low enough to heal")
											time.Sleep(1 * time.Second)
										} else if hp1 >= (hp1or - attack3){
											time.Sleep(1 * time.Second)
											hp1 = hp1or
											fmt.Println("You have been healed to max")
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
									time.Sleep(1 * time.Second)
									fmt.Println("enemy is attacking")
									time.Sleep(1 * time.Second)
									if hp2 <= (hp2or / 2) && hp1 > attacken2 {
										time.Sleep(1 * time.Second)
										hp2 = hp2 + attacken3
										fmt.Println("enemy has been healed",attacken3,"health")
										time.Sleep(1 * time.Second)
										continue
									}
									if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp1 = hp1 - attacken1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}	
										}else {
											fmt.Println("you took",attacken1,"damage")
											hp1 = hp1-attacken1
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}
										}
									} else {
										var b = time.Now().UnixNano()
										rand.Seed(b)
										var random3 = rand.Intn(5)
											if random3 == 1 || random3 == 2 || random3 == 3 {
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 0 {
													hp1 = hp1 - attacken2 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
													time.Sleep(1 * time.Second)
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2
														break
													}	
												}else {
													fmt.Println("you took",attacken2,"damage")
													hp1 = hp1-attacken2
													time.Sleep(1 * time.Second)	
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2
														break
													}
												}
											} else {
												fmt.Println("attack missed")
													time.Sleep(1 * time.Second)	
													continue
											}
										}
									}
									if check == 2 {
										time.Sleep(2 * time.Second)
										fmt.Println("Respawning")
										time.Sleep(5 * time.Second)
										key = 1
										continue
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
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack4 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
												hp2 = hp2 - attack4
												fmt.Println("Enemy took",attack4,"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
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
									fmt.Println()
									fmt.Println("choose your attack")
									fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
									fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
									fmt.Println("3.heal",attack3,"health 100% accuracy")
									fmt.Scanln(&a)
									if a == 1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp2 = hp2 - attack1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1,"damage")
											hp2 = hp2-attack1
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
										}
									} else if a == 2 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(3)
										if random == 1 || random == 2 || random == 3 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack2 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
											time.Sleep(1 * time.Second)
											hp2 = hp2 - attack2
											fmt.Println("Enemy took",attack2,"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
											}
										} else {
											time.Sleep(1 * time.Second)
											fmt.Println("attack missed")
											time.Sleep(1 * time.Second)
										}
									} else if a == 3 {
										if hp1 == hp1or {
											time.Sleep(1 * time.Second)
											fmt.Println("You're not low enough to heal")
											time.Sleep(1 * time.Second)
										} else if hp1 >= (hp1or - attack3){
											time.Sleep(1 * time.Second)
											hp1 = hp1or
											fmt.Println("You have been healed to max")
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
									time.Sleep(1 * time.Second)
									fmt.Println("enemy is attacking")
									time.Sleep(1 * time.Second)
									if hp2 <= (hp2or / 2) && hp1 > attacken2 {
										time.Sleep(1 * time.Second)
										hp2 = hp2 + attacken3
										fmt.Println("enemy has been healed",attacken3,"health")
										time.Sleep(1 * time.Second)
										continue
									}
									if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp1 = hp1 - attacken1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}	
										}else {
											fmt.Println("you took",attacken1,"damage")
											hp1 = hp1-attacken1
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}
										}
									} else {
										var b = time.Now().UnixNano()
										rand.Seed(b)
										var random3 = rand.Intn(5)
											if random3 == 1 || random3 == 2 || random3 == 3 {
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 0 {
													hp1 = hp1 - attacken2 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
													time.Sleep(1 * time.Second)
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2
														break
													}	
												}else {
													fmt.Println("you took",attacken2,"damage")
													hp1 = hp1-attacken2
													time.Sleep(1 * time.Second)	
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2
														break
													}
												}
											} else {
												fmt.Println("attack missed")
													time.Sleep(1 * time.Second)	
													continue
											}
										}
									}
								if check == 2 {
									time.Sleep(2 * time.Second)
									fmt.Println("Respawning")
									time.Sleep(5 * time.Second)
									key2 = 1
									continue
								}
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
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack4 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
												hp2 = hp2 - attack4
												fmt.Println("Enemy took",attack4,"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
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
									fmt.Println()
									fmt.Println("choose your attack")
									fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
									fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
									fmt.Println("3.heal",attack3,"health 100% accuracy")
									fmt.Scanln(&a)
									if a == 1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp2 = hp2 - attack1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1,"damage")
											hp2 = hp2-attack1
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
										}
									} else if a == 2 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(3)
										if random == 1 || random == 2 || random == 3 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack2 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
											time.Sleep(1 * time.Second)
											hp2 = hp2 - attack2
											fmt.Println("Enemy took",attack2,"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
											}
										} else {
											time.Sleep(1 * time.Second)
											fmt.Println("attack missed")
											time.Sleep(1 * time.Second)
										}
									} else if a == 3 {
										if hp1 == hp1or {
											time.Sleep(1 * time.Second)
											fmt.Println("You're not low enough to heal")
											time.Sleep(1 * time.Second)
										} else if hp1 >= (hp1or - attack3){
											time.Sleep(1 * time.Second)
											hp1 = hp1or
											fmt.Println("You have been healed to max")
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
									time.Sleep(1 * time.Second)
									fmt.Println("enemy is attacking")
									time.Sleep(1 * time.Second)
									if hp2 <= (hp2or / 2) && hp1 > attacken2 {
										time.Sleep(1 * time.Second)
										hp2 = hp2 + attacken3
										fmt.Println("enemy has been healed",attacken3,"health")
										time.Sleep(1 * time.Second)
										continue
									}
									if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp1 = hp1 - attacken1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}	
										}else {
											fmt.Println("you took",attacken1,"damage")
											hp1 = hp1-attacken1
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}
										}
									} else {
										var b = time.Now().UnixNano()
										rand.Seed(b)
										var random3 = rand.Intn(5)
											if random3 == 1 || random3 == 2 || random3 == 3 {
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 0 {
													hp1 = hp1 - attacken2 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
													time.Sleep(1 * time.Second)
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2
														break
													}	
												}else {
													fmt.Println("you took",attacken2,"damage")
													hp1 = hp1-attacken2
													time.Sleep(1 * time.Second)	
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2
														break
													}
												}
											} else {
												fmt.Println("attack missed")
													time.Sleep(1 * time.Second)	
													continue
											}
										}
									}
								if check == 2 {
									time.Sleep(2 * time.Second)
									fmt.Println("Respawning")
									time.Sleep(5 * time.Second)
									key3 = 1
									continue
								}
								hp1 = hp1or
								mon = mon + 50
								time.Sleep(1 * time.Second)
							fmt.Println("You gained 50 coins")
							}else{
								fmt.Println("You already explored here. departing in 3 seconds")
								time.Sleep(3 * time.Second)
							}
						}else if a == 4 {
							time.Sleep(1 * time.Second)
								fmt.Println("You enter a dungeon with a large monster weilding a giant axe")
								time.Sleep(1 * time.Second)
								fmt.Println("You enter a battle with Taurus")
								hp2 = hp2or
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
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack4 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
												hp2 = hp2 - attack4
												fmt.Println("Enemy took",attack4,"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
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
									fmt.Println()
									fmt.Println("choose your attack")
									fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
									fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
									fmt.Println("3.heal",attack3,"health 100% accuracy")
									fmt.Scanln(&a)
									if a == 1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp2 = hp2 - attack1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1,"damage")
											hp2 = hp2-attack1
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
										}
									} else if a == 2 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(3)
										if random == 1 || random == 2 || random == 3 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack2 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
											time.Sleep(1 * time.Second)
											hp2 = hp2 - attack2
											fmt.Println("Enemy took",attack2,"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
											}
										} else {
											time.Sleep(1 * time.Second)
											fmt.Println("attack missed")
											time.Sleep(1 * time.Second)
										}
									} else if a == 3 {
										if hp1 == hp1or {
											time.Sleep(1 * time.Second)
											fmt.Println("You're not low enough to heal")
											time.Sleep(1 * time.Second)
										} else if hp1 >= (hp1or - attack3){
											time.Sleep(1 * time.Second)
											hp1 = hp1or
											fmt.Println("You have been healed to max")
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
									time.Sleep(1 * time.Second)
									fmt.Println("enemy is attacking")
									time.Sleep(1 * time.Second)
									if hp2 <= (hp2or / 2) && hp1 > attacken2 {
										time.Sleep(1 * time.Second)
										hp2 = hp2 + attacken3
										fmt.Println("enemy has been healed",attacken3,"health")
										time.Sleep(1 * time.Second)
										continue
									}
									if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp1 = hp1 - attacken1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}	
										}else {
											fmt.Println("you took",attacken1,"damage")
											hp1 = hp1-attacken1
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}
										}
									} else {
										var b = time.Now().UnixNano()
										rand.Seed(b)
										var random3 = rand.Intn(5)
											if random3 == 1 || random3 == 2 || random3 == 3 {
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 0 {
													hp1 = hp1 - attacken2 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
													time.Sleep(1 * time.Second)
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2
														break
													}	
												}else {
													fmt.Println("you took",attacken2,"damage")
													hp1 = hp1-attacken2
													time.Sleep(1 * time.Second)	
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2
														break
													}
												}
											} else {
												fmt.Println("attack missed")
													time.Sleep(1 * time.Second)	
													continue
											}
										}
									}
								if check == 2 {
									time.Sleep(2 * time.Second)
									fmt.Println("Respawning")
									time.Sleep(5 * time.Second)
									continue
								}
								hp1 = hp1or
								time.Sleep(4 * time.Second)
								mon = mon + 200
								time.Sleep(1 * time.Second)
							fmt.Println("You gained 200 coins")
							time.Sleep(1 * time.Second)
							fmt.Println("Total Coins:",mon)
							time.Sleep(1 * time.Second)
								fmt.Println("After beating the Taurus, you scavenge its carcass, and find a key, which opens the chest.")
								time.Sleep(4 * time.Second)
								fmt.Println("The first, bright yellow gem is inside. You have saved Mystic Woods from the curse.")
								time.Sleep(4 * time.Second)
								fmt.Println("Now, you must go, and conquer Mount Olympus")
							time.Sleep(7 * time.Second)
							check4 = 1
							break
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
								if check == 2 {
							time.Sleep(2 * time.Second)
							fmt.Println("Respawning")
							time.Sleep(5 * time.Second)
							continue
						}
						if a >= 4 || a == 0{
							continue
						}
						if check4 == 1 {
							break
						}
						attack1tr = 0
						hp1tr = 0
						attack2tr = 0
						attack3tr = 0
						attack4tr = 0
						montr = 0
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
				for {
				hp1 = hp1 - hp1tr 
				hp1 = hp1or
				attack1 = attack1 - attack1tr
				mon = mon - montr
				hp1 = hp1 - hp1tr
				hp1or = hp1
				attack2 = attack2 - attack2tr
				attack3 = attack3 - attack3tr
				attack4 = attack4 - attack4tr
				hp2or = hp2
				attack1tr = 0
				hp1tr = 0
				attack2tr = 0
				attack3tr = 0
				attack4tr = 0
				montr = 0
				check = 1
					attacken1 = 20
					attacken2 = 40
					hp2 = 130
					hp2or = hp2
					attacken3 = 30
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
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 3 {
									hp2 = hp2 - attack4 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}else {
									hp2 = hp2 - attack4
									fmt.Println("Enemy took",attack4,"damage")
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}
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
						fmt.Println()
						fmt.Println("choose your attack")
						fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
						fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
						fmt.Println("3.heal",attack3,"health 100% accuracy")
						fmt.Scanln(&a)
						if a == 1 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(8)
							if random == 0 {
								hp2 = hp2 - attack1 * 3 / 2
								fmt.Println("Critical Hit")
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}	
							}else {
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy took",attack1,"damage")
								hp2 = hp2-attack1
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}
							}
						} else if a == 2 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(3)
							if random == 1 || random == 2 || random == 3 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 3 {
									hp2 = hp2 - attack2 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}else {
								time.Sleep(1 * time.Second)
								hp2 = hp2 - attack2
								fmt.Println("Enemy took",attack2,"damage")
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}
								}
							} else {
								time.Sleep(1 * time.Second)
								fmt.Println("attack missed")
								time.Sleep(1 * time.Second)
							}
						} else if a == 3 {
							if hp1 == hp1or {
								time.Sleep(1 * time.Second)
								fmt.Println("You're not low enough to heal")
								time.Sleep(1 * time.Second)
							} else if hp1 >= (hp1or - attack3){
								time.Sleep(1 * time.Second)
								hp1 = hp1or
								fmt.Println("You have been healed to max")
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
						time.Sleep(1 * time.Second)
						fmt.Println("enemy is attacking")
						time.Sleep(1 * time.Second)
						if hp2 <= (hp2or / 2) && hp1 > attacken2 {
							time.Sleep(1 * time.Second)
							hp2 = hp2 + attacken3
							fmt.Println("enemy has been healed",attacken3,"health")
							time.Sleep(1 * time.Second)
							continue
						}
						if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(8)
							if random == 0 {
								hp1 = hp1 - attacken1 * 3 / 2
								fmt.Println("Critical Hit")
								time.Sleep(1 * time.Second)
								fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
								time.Sleep(1 * time.Second)
								if hp1 <= 0 {
									fmt.Println("You Died")
									check = 2
									break
								}	
							}else {
								fmt.Println("you took",attacken1,"damage")
								hp1 = hp1-attacken1
								time.Sleep(1 * time.Second)
								if hp1 <= 0 {
									fmt.Println("You Died")
									check = 2
									break
								}
							}
						} else {
							var b = time.Now().UnixNano()
							rand.Seed(b)
							var random3 = rand.Intn(5)
								if random3 == 1 || random3 == 2 || random3 == 3 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp1 = hp1 - attacken2 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											check = 2
											break
										}	
									}else {
										fmt.Println("you took",attacken2,"damage")
										hp1 = hp1-attacken2
										time.Sleep(1 * time.Second)	
										if hp1 <= 0 {
											fmt.Println("You Died")
											check = 2
											break
										}
									}
								} else {
									fmt.Println("attack missed")
										time.Sleep(1 * time.Second)	
										continue
								}
							}
						}
					if check == 2 {
						time.Sleep(2 * time.Second)
						fmt.Println("Respawning")
						time.Sleep(5 * time.Second)
						continue
					}
					hp1 = hp1or
					hp1 = hp1 + 30
					hp1tr = hp1tr + 30
					hp1or = hp1
					time.Sleep(1 * time.Second)
					fmt.Println("After the battle, you claim its helmet.")
					time.Sleep(4 * time.Second)
					fmt.Println("Armour upgraded")
					time.Sleep(4 * time.Second)
					fmt.Println("You know what's coming for you. There is no good left in this Mountain range. You are aware that you must conquer two other mountains before Olympus.")
					time.Sleep(4 * time.Second)
					fmt.Println("The first mountain, is merely  just a hill. On the top of the mountain, waits the Brood Mother, a lethally dangerous spider.")
					time.Sleep(4 * time.Second)
					fmt.Println("When you hike up the hill, you notice some movement. It is a possessed spartan.  As soon as you notice it, the spartan gets impaled by a ten-foot spider.")
					time.Sleep(4 * time.Second)
					fmt.Println("You run away into a bush, hiding. The spider notices you.")
					hp2 = 190
					hp2or = hp2
					attacken1 = 30
					attacken2 = 50
					attacken3 = 30
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
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 3 {
									hp2 = hp2 - attack4 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}else {
									hp2 = hp2 - attack4
									fmt.Println("Enemy took",attack4,"damage")
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}
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
						fmt.Println()
						fmt.Println("choose your attack")
						fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
						fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
						fmt.Println("3.heal",attack3,"health 100% accuracy")
						fmt.Scanln(&a)
						if a == 1 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(8)
							if random == 0 {
								hp2 = hp2 - attack1 * 3 / 2
								fmt.Println("Critical Hit")
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}	
							}else {
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy took",attack1,"damage")
								hp2 = hp2-attack1
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}
							}
						} else if a == 2 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(3)
							if random == 1 || random == 2 || random == 3 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 3 {
									hp2 = hp2 - attack2 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}else {
								time.Sleep(1 * time.Second)
								hp2 = hp2 - attack2
								fmt.Println("Enemy took",attack2,"damage")
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}
								}
							} else {
								time.Sleep(1 * time.Second)
								fmt.Println("attack missed")
								time.Sleep(1 * time.Second)
							}
						} else if a == 3 {
							if hp1 == hp1or {
								time.Sleep(1 * time.Second)
								fmt.Println("You're not low enough to heal")
								time.Sleep(1 * time.Second)
							} else if hp1 >= (hp1or - attack3){
								time.Sleep(1 * time.Second)
								hp1 = hp1or
								fmt.Println("You have been healed to max")
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
						time.Sleep(1 * time.Second)
						fmt.Println("enemy is attacking")
						time.Sleep(1 * time.Second)
						if hp2 <= (hp2or / 2) && hp1 > attacken2 {
							time.Sleep(1 * time.Second)
							hp2 = hp2 + attacken3
							fmt.Println("enemy has been healed",attacken3,"health")
							time.Sleep(1 * time.Second)
							continue
						}
						if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(8)
							if random == 0 {
								hp1 = hp1 - attacken1 * 3 / 2
								fmt.Println("Critical Hit")
								time.Sleep(1 * time.Second)
								fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
								time.Sleep(1 * time.Second)
								if hp1 <= 0 {
									fmt.Println("You Died")
									check = 2
									break
								}	
							}else {
								fmt.Println("you took",attacken1,"damage")
								hp1 = hp1-attacken1
								time.Sleep(1 * time.Second)
								if hp1 <= 0 {
									fmt.Println("You Died")
									check = 2
									break
								}
							}
						} else {
							var b = time.Now().UnixNano()
							rand.Seed(b)
							var random3 = rand.Intn(5)
								if random3 == 1 || random3 == 2 || random3 == 3 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp1 = hp1 - attacken2 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											check = 2
											break
										}	
									}else {
										fmt.Println("you took",attacken2,"damage")
										hp1 = hp1-attacken2
										time.Sleep(1 * time.Second)	
										if hp1 <= 0 {
											fmt.Println("You Died")
											check = 2
											break
										}
									}
								} else {
									fmt.Println("attack missed")
										time.Sleep(1 * time.Second)	
										continue
								}
							}
						}
						if check == 2 {
							time.Sleep(2 * time.Second)
							fmt.Println("Respawning")
							time.Sleep(5 * time.Second)
							continue
						}
						attack1tr = 0
						hp1tr = 0
						attack2tr = 0
						attack3tr = 0
						attack4tr = 0
						montr = 0
						break
				}
				for {
				hp1 = hp1 - hp1tr 
				hp1 = hp1or
				attack1 = attack1 - attack1tr
				mon = mon - montr
				hp1 = hp1 - hp1tr
				hp1or = hp1
				attack2 = attack2 - attack2tr
				attack3 = attack3 - attack3tr
				attack4 = attack4 - attack4tr
				hp2or = hp2
				attack1tr = 0
				hp1tr = 0
				attack2tr = 0
				attack3tr = 0
				attack4tr = 0
				montr = 0
				check = 1
				time.Sleep(4 * time.Second)
				fmt.Println("The Brood Mother flees, and you continue up the hill. You hear a creak in the distance, Do you go there:")
				time.Sleep(4 * time.Second)
				fmt.Println("1.Yes")
				fmt.Println("2.No")
				fmt.Scanln(&b)
				if b == 1 {
					fmt.Println("You descend the hill, but you trip, fall and die.")
					time.Sleep(4 * time.Second)
					fmt.Println("Respawning")
					time.Sleep(5 * time.Second)
					continue
				}
				fmt.Println("You have just avoided a very disappointing death. You nervously advance forward.")
				time.Sleep(4 * time.Second)
				fmt.Println("you have reached the top of the hill. but you don't see the spider. What do you do? Ring the bell, or run back")
				time.Sleep(4 * time.Second)
				fmt.Println("1.Ring bell")
				fmt.Println("2.Run back")
				fmt.Scanln(&b)
				if b == 2 {
					fmt.Println("You run down the hill, scared. you feel a sharp pain in the back.")
					time.Sleep(4 * time.Second)
					fmt.Println("Just like the possessed Spartan, you have been impaled. Dead.")
					time.Sleep( 4 * time.Second)
					fmt.Println("Respawning")
					time.Sleep( 5 * time.Second)
					continue
				}
				time.Sleep(4 * time.Second)
				fmt.Println("You ring the bell, you have effortlessly conquered the first Mountain.")
				time.Sleep(4 * time.Second)
				fmt.Println("You advance towards the second Mountain, which is way larger than the first, but way smaller than Mt. Olympus.")
				time.Sleep(4 * time.Second)
				fmt.Println("You are stopped in your path. It is the Brood Mother.")
				hp1 = hp1or
				hp2 = 100
				hp2or = hp2
				attacken1 = 20
				attacken2 = 40
				attacken3 = 30
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
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(8)
							if random == 3 {
								hp2 = hp2 - attack4 * 3 / 2
								fmt.Println("Critical Hit")
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}	
							}else {
								hp2 = hp2 - attack4
								fmt.Println("Enemy took",attack4,"damage")
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}
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
					fmt.Println()
					fmt.Println("choose your attack")
					fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
					fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
					fmt.Println("3.heal",attack3,"health 100% accuracy")
					fmt.Scanln(&a)
					if a == 1 {
						var s = time.Now().UnixNano()
						rand.Seed(s)
						var random = rand.Intn(8)
						if random == 0 {
							hp2 = hp2 - attack1 * 3 / 2
							fmt.Println("Critical Hit")
							time.Sleep(1 * time.Second)
							fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
							time.Sleep(1 * time.Second)
							if hp2 <= 0 {
								fmt.Println("You Win")
								break
							}	
						}else {
							time.Sleep(1 * time.Second)
							fmt.Println("Enemy took",attack1,"damage")
							hp2 = hp2-attack1
							time.Sleep(1 * time.Second)
							if hp2 <= 0 {
								fmt.Println("You Win")
								break
							}
						}
					} else if a == 2 {
						var s = time.Now().UnixNano()
						rand.Seed(s)
						var random = rand.Intn(3)
						if random == 1 || random == 2 || random == 3 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(8)
							if random == 3 {
								hp2 = hp2 - attack2 * 3 / 2
								fmt.Println("Critical Hit")
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}	
							}else {
							time.Sleep(1 * time.Second)
							hp2 = hp2 - attack2
							fmt.Println("Enemy took",attack2,"damage")
							time.Sleep(1 * time.Second)
							if hp2 <= 0 {
								fmt.Println("You Win")
								break
							}
							}
						} else {
							time.Sleep(1 * time.Second)
							fmt.Println("attack missed")
							time.Sleep(1 * time.Second)
						}
					} else if a == 3 {
						if hp1 == hp1or {
							time.Sleep(1 * time.Second)
							fmt.Println("You're not low enough to heal")
							time.Sleep(1 * time.Second)
						} else if hp1 >= (hp1or - attack3){
							time.Sleep(1 * time.Second)
							hp1 = hp1or
							fmt.Println("You have been healed to max")
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
					time.Sleep(1 * time.Second)
					fmt.Println("enemy is attacking")
					time.Sleep(1 * time.Second)
					if hp2 <= (hp2or / 2) && hp1 > attacken2 {
						time.Sleep(1 * time.Second)
						hp2 = hp2 + attacken3
						fmt.Println("enemy has been healed",attacken3,"health")
						time.Sleep(1 * time.Second)
						continue
					}
					if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
						var s = time.Now().UnixNano()
						rand.Seed(s)
						var random = rand.Intn(8)
						if random == 0 {
							hp1 = hp1 - attacken1 * 3 / 2
							fmt.Println("Critical Hit")
							time.Sleep(1 * time.Second)
							fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
							time.Sleep(1 * time.Second)
							if hp1 <= 0 {
								fmt.Println("You Died")
								check = 2
								break
							}	
						}else {
							fmt.Println("you took",attacken1,"damage")
							hp1 = hp1-attacken1
							time.Sleep(1 * time.Second)
							if hp1 <= 0 {
								fmt.Println("You Died")
								check = 2
								break
							}
						}
					} else {
						var b = time.Now().UnixNano()
						rand.Seed(b)
						var random3 = rand.Intn(5)
							if random3 == 1 || random3 == 2 || random3 == 3 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 0 {
									hp1 = hp1 - attacken2 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
									time.Sleep(1 * time.Second)
									if hp1 <= 0 {
										fmt.Println("You Died")
										check = 2
										break
									}	
								}else {
									fmt.Println("you took",attacken2,"damage")
									hp1 = hp1-attacken2
									time.Sleep(1 * time.Second)	
									if hp1 <= 0 {
										fmt.Println("You Died")
										check = 2
										break
									}
								}
							} else {
								fmt.Println("attack missed")
									time.Sleep(1 * time.Second)	
									continue
							}
						}
					}
					if check == 2 {
						time.Sleep(2 * time.Second)
						fmt.Println("Respawning")
						time.Sleep(5 * time.Second)
						continue
					}
				hp1 = hp1or
				attack1tr = 0
				hp1tr = 0
				attack2tr = 0
				attack3tr = 0
				attack4tr = 0
				montr = 0
				break
				}
						attack3 = attack3 + 20
						fmt.Println("The Brood Mother was powerless. You had reduced her to cripples in the first battle.")
						time.Sleep(4 * time.Second)
						fmt.Println("You scavenge her body, and find an advanced heal potion.")
						time.Sleep(4 * time.Second)
						fmt.Println("Magic upgraded")
						time.Sleep(4 * time.Second)
						fmt.Println("The second mountain, home of the Mountain beast, who has low health, but packs a great punch.")
						time.Sleep(4 * time.Second)
						fmt.Println("He guards the second bell, so you advance up the mountain. on the path, you encounter a Lost Triton.")
						time.Sleep(4 * time.Second)
						fmt.Println("A minion which wields a trident.")
						for {
							hp1 = hp1 - hp1tr 
							hp1 = hp1or
							attack1 = attack1 - attack1tr
							mon = mon - montr
							hp1 = hp1 - hp1tr
							hp1or = hp1
							attack2 = attack2 - attack2tr
							attack3 = attack3 - attack3tr
							attack4 = attack4 - attack4tr
							hp2or = hp2
							attack1tr = 0
							hp1tr = 0
							attack2tr = 0
							attack3tr = 0
							attack4tr = 0
							montr = 0
							check = 1
							hp2 = 130
							hp2or = hp2
							attacken1 = 35
							attacken2 = 55
							attacken3 = 30
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
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 3 {
											hp2 = hp2 - attack4 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
											hp2 = hp2 - attack4
											fmt.Println("Enemy took",attack4,"damage")
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
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
								fmt.Println()
								fmt.Println("choose your attack")
								fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
								fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
								fmt.Println("3.heal",attack3,"health 100% accuracy")
								fmt.Scanln(&a)
								if a == 1 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp2 = hp2 - attack1 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									}else {
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack1,"damage")
										hp2 = hp2-attack1
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
									}
								} else if a == 2 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(3)
									if random == 1 || random == 2 || random == 3 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 3 {
											hp2 = hp2 - attack2 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
										time.Sleep(1 * time.Second)
										hp2 = hp2 - attack2
										fmt.Println("Enemy took",attack2,"damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
										}
									} else {
										time.Sleep(1 * time.Second)
										fmt.Println("attack missed")
										time.Sleep(1 * time.Second)
									}
								} else if a == 3 {
									if hp1 == hp1or {
										time.Sleep(1 * time.Second)
										fmt.Println("You're not low enough to heal")
										time.Sleep(1 * time.Second)
									} else if hp1 >= (hp1or - attack3){
										time.Sleep(1 * time.Second)
										hp1 = hp1or
										fmt.Println("You have been healed to max")
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
								time.Sleep(1 * time.Second)
								fmt.Println("enemy is attacking")
								time.Sleep(1 * time.Second)
								if hp2 <= (hp2or / 2) && hp1 > attacken2 {
									time.Sleep(1 * time.Second)
									hp2 = hp2 + attacken3
									fmt.Println("enemy has been healed",attacken3,"health")
									time.Sleep(1 * time.Second)
									continue
								}
								if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp1 = hp1 - attacken1 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											check = 2
											break
										}	
									}else {
										fmt.Println("you took",attacken1,"damage")
										hp1 = hp1-attacken1
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											check = 2
											break
										}
									}
								} else {
									var b = time.Now().UnixNano()
									rand.Seed(b)
									var random3 = rand.Intn(5)
										if random3 == 1 || random3 == 2 || random3 == 3 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp1 = hp1 - attacken2 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2
													break
												}	
											}else {
												fmt.Println("you took",attacken2,"damage")
												hp1 = hp1-attacken2
												time.Sleep(1 * time.Second)	
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2
													break
												}
											}
										} else {
											fmt.Println("attack missed")
												time.Sleep(1 * time.Second)	
												continue
										}
									}
								}
								if check == 2 {
									time.Sleep(2 * time.Second)
									fmt.Println("Respawning")
									time.Sleep(5 * time.Second)
									continue
								}
							hp1 = hp1or
							time.Sleep(4 * time.Second)
							fmt.Println("After beating the Triton, you notice it dropped something, do you ignore it or pick it up ?")
							time.Sleep(4 * time.Second)
							fmt.Println("1.Pick it up.")
							fmt.Println("2.Ignore it.")
							fmt.Scanln(&b)
							if b == 1 {
								fmt.Println("New Weapon: Trident")
								attack1tr = attack1
								attack2tr = attack2
								attack4tr = attack4
								attack2 = 75
								attack1 = 50
								attack4 = 100
								attack1tr = attack1 - attack1tr
								attack2tr = attack2 - attack2tr
								attack4tr = attack4 - attack4tr
							}
							time.Sleep(4 * time.Second)
							fmt.Println("you proceed further up, until you see a wall, too hard to climb.")
							time.Sleep(4 * time.Second)
							fmt.Println("You attempt to climb the wall, and notice that you are at a plateau, almost up the mountain. You encounter a demon.")
							time.Sleep(4 * time.Second)
							fmt.Println("It is one of Hades' minions.")
							time.Sleep(4 * time.Second)
							hp2 = 180
							hp2or = hp2
							attacken1 = 35
							attacken2 = 65
							attacken3 = 30
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
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 3 {
											hp2 = hp2 - attack4 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
											hp2 = hp2 - attack4
											fmt.Println("Enemy took",attack4,"damage")
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
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
								fmt.Println()
								fmt.Println("choose your attack")
								fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
								fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
								fmt.Println("3.heal",attack3,"health 100% accuracy")
								fmt.Scanln(&a)
								if a == 1 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp2 = hp2 - attack1 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									}else {
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack1,"damage")
										hp2 = hp2-attack1
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
									}
								} else if a == 2 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(3)
									if random == 1 || random == 2 || random == 3 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 3 {
											hp2 = hp2 - attack2 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
										time.Sleep(1 * time.Second)
										hp2 = hp2 - attack2
										fmt.Println("Enemy took",attack2,"damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
										}
									} else {
										time.Sleep(1 * time.Second)
										fmt.Println("attack missed")
										time.Sleep(1 * time.Second)
									}
								} else if a == 3 {
									if hp1 == hp1or {
										time.Sleep(1 * time.Second)
										fmt.Println("You're not low enough to heal")
										time.Sleep(1 * time.Second)
									} else if hp1 >= (hp1or - attack3){
										time.Sleep(1 * time.Second)
										hp1 = hp1or
										fmt.Println("You have been healed to max")
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
								time.Sleep(1 * time.Second)
								fmt.Println("enemy is attacking")
								time.Sleep(1 * time.Second)
								if hp2 <= (hp2or / 2) && hp1 > attacken2 {
									time.Sleep(1 * time.Second)
									hp2 = hp2 + attacken3
									fmt.Println("enemy has been healed",attacken3,"health")
									time.Sleep(1 * time.Second)
									continue
								}
								if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp1 = hp1 - attacken1 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											check = 2
											break
										}	
									}else {
										fmt.Println("you took",attacken1,"damage")
										hp1 = hp1-attacken1
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											check = 2
											break
										}
									}
								} else {
									var b = time.Now().UnixNano()
									rand.Seed(b)
									var random3 = rand.Intn(5)
										if random3 == 1 || random3 == 2 || random3 == 3 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp1 = hp1 - attacken2 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2
													break
												}	
											}else {
												fmt.Println("you took",attacken2,"damage")
												hp1 = hp1-attacken2
												time.Sleep(1 * time.Second)	
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2
													break
												}
											}
										} else {
											fmt.Println("attack missed")
												time.Sleep(1 * time.Second)	
												continue
										}
									}
								}
								if check == 2 {
									time.Sleep(2 * time.Second)
									fmt.Println("Respawning")
									time.Sleep(5 * time.Second)
									continue
								}
							hp1 = hp1or
							attack1tr = 0
							hp1tr = 0
							attack2tr = 0
							attack3tr = 0
							attack4tr = 0
							montr = 0
							break
						}
					hp1 = hp1 + 20
					hp1or = hp1
					for {
						fmt.Println("After beating the minion, your head feels a little relieved.")
						time.Sleep(4 * time.Second)
						fmt.Println("You pick up his helmet")
						time.Sleep(4 * time.Second)
						fmt.Println("Armour upgraded")
						time.Sleep(4 * time.Second)
						fmt.Println("This means that if you get rid of their master, it undo's the curse. You continue up the hill, and reach the top.")
						time.Sleep(4 * time.Second)
						fmt.Println("You see a Skeletal Giant. You must beat him to ring the second bell. You confidently enter the battle.")
						time.Sleep(4 * time.Second)
						hp1 = hp1 - hp1tr 
						hp1 = hp1or
						attack1 = attack1 - attack1tr
						mon = mon - montr
						hp1 = hp1 - hp1tr
						hp1or = hp1
						attack2 = attack2 - attack2tr
						attack3 = attack3 - attack3tr
						attack4 = attack4 - attack4tr
						hp2or = hp2
						attack1tr = 0
						hp1tr = 0
						attack2tr = 0
						attack3tr = 0
						attack4tr = 0
						montr = 0
						check = 1
						hp2 = 210
						hp2or = hp2
						attacken1 = 40
						attacken2 = 80
						attacken3 = 30
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
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 3 {
										hp2 = hp2 - attack4 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									}else {
										hp2 = hp2 - attack4
										fmt.Println("Enemy took",attack4,"damage")
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
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
							fmt.Println()
							fmt.Println("choose your attack")
							fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
							fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
							fmt.Println("3.heal",attack3,"health 100% accuracy")
							fmt.Scanln(&a)
							if a == 1 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 0 {
									hp2 = hp2 - attack1 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}else {
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack1,"damage")
									hp2 = hp2-attack1
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}
								}
							} else if a == 2 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(3)
								if random == 1 || random == 2 || random == 3 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 3 {
										hp2 = hp2 - attack2 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									}else {
									time.Sleep(1 * time.Second)
									hp2 = hp2 - attack2
									fmt.Println("Enemy took",attack2,"damage")
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}
									}
								} else {
									time.Sleep(1 * time.Second)
									fmt.Println("attack missed")
									time.Sleep(1 * time.Second)
								}
							} else if a == 3 {
								if hp1 == hp1or {
									time.Sleep(1 * time.Second)
									fmt.Println("You're not low enough to heal")
									time.Sleep(1 * time.Second)
								} else if hp1 >= (hp1or - attack3){
									time.Sleep(1 * time.Second)
									hp1 = hp1or
									fmt.Println("You have been healed to max")
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
							time.Sleep(1 * time.Second)
							fmt.Println("enemy is attacking")
							time.Sleep(1 * time.Second)
							if hp2 <= (hp2or / 2) && hp1 > attacken2 {
								time.Sleep(1 * time.Second)
								hp2 = hp2 + attacken3
								fmt.Println("enemy has been healed",attacken3,"health")
								time.Sleep(1 * time.Second)
								continue
							}
							if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 0 {
									hp1 = hp1 - attacken1 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
									time.Sleep(1 * time.Second)
									if hp1 <= 0 {
										fmt.Println("You Died")
										check = 2
										break
									}	
								}else {
									fmt.Println("you took",attacken1,"damage")
									hp1 = hp1-attacken1
									time.Sleep(1 * time.Second)
									if hp1 <= 0 {
										fmt.Println("You Died")
										check = 2
										break
									}
								}
							} else {
								var b = time.Now().UnixNano()
								rand.Seed(b)
								var random3 = rand.Intn(5)
									if random3 == 1 || random3 == 2 || random3 == 3 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp1 = hp1 - attacken2 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}	
										}else {
											fmt.Println("you took",attacken2,"damage")
											hp1 = hp1-attacken2
											time.Sleep(1 * time.Second)	
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}
										}
									} else {
										fmt.Println("attack missed")
											time.Sleep(1 * time.Second)	
											continue
									}
								}
							}
							if check == 2 {
								time.Sleep(2 * time.Second)
								fmt.Println("Respawning")
								time.Sleep(5 * time.Second)
								continue
							}
						time.Sleep(4 * time.Second)
						fmt.Println("You have defeated the Skeletal Giant. you head over to the second bell, and ring it.")
						time.Sleep(4 * time.Second)
						fmt.Println("You see his hammer on the ground do you pick it up.")
						fmt.Println("1.Pick it up.")
						fmt.Println("2.Ignore it")
						fmt.Scanln(&b)
						if b == 1 {
							attack1tr = attack1
							attack2tr = attack2
							attack4tr = attack4
							hp1 = hp1or
							attack1 = 60
							attack2 = 80
							attack4 = 120
							attack1tr = attack1 - attack1tr
							attack2tr = attack2 - attack2tr
							attack4tr = attack4 - attack4tr
							time.Sleep(4 * time.Second)
							fmt.Println("Weapons upgraded")
							time.Sleep(4 * time.Second)
						}
						fmt.Println("You can now advance to Mount Olympus. The largest mountain, overlooking Mystic Woods.")
						time.Sleep(4 * time.Second)
						fmt.Println("You begin an almost endless trek, and when you are about halfway there, you encounter a possessed spartan. ")
						time.Sleep(4 * time.Second)
						hp2 = 140
						hp2or = hp2
						attacken1 = 30
						attacken2 = 50
						attacken3 = 30
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
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 3 {
										hp2 = hp2 - attack4 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									}else {
										hp2 = hp2 - attack4
										fmt.Println("Enemy took",attack4,"damage")
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
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
							fmt.Println()
							fmt.Println("choose your attack")
							fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
							fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
							fmt.Println("3.heal",attack3,"health 100% accuracy")
							fmt.Scanln(&a)
							if a == 1 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 0 {
									hp2 = hp2 - attack1 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}else {
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack1,"damage")
									hp2 = hp2-attack1
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}
								}
							} else if a == 2 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(3)
								if random == 1 || random == 2 || random == 3 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 3 {
										hp2 = hp2 - attack2 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									}else {
									time.Sleep(1 * time.Second)
									hp2 = hp2 - attack2
									fmt.Println("Enemy took",attack2,"damage")
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}
									}
								} else {
									time.Sleep(1 * time.Second)
									fmt.Println("attack missed")
									time.Sleep(1 * time.Second)
								}
							} else if a == 3 {
								if hp1 == hp1or {
									time.Sleep(1 * time.Second)
									fmt.Println("You're not low enough to heal")
									time.Sleep(1 * time.Second)
								} else if hp1 >= (hp1or - attack3){
									time.Sleep(1 * time.Second)
									hp1 = hp1or
									fmt.Println("You have been healed to max")
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
							time.Sleep(1 * time.Second)
							fmt.Println("enemy is attacking")
							time.Sleep(1 * time.Second)
							if hp2 <= (hp2or / 2) && hp1 > attacken2 {
								time.Sleep(1 * time.Second)
								hp2 = hp2 + attacken3
								fmt.Println("enemy has been healed",attacken3,"health")
								time.Sleep(1 * time.Second)
								continue
							}
							if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 0 {
									hp1 = hp1 - attacken1 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
									time.Sleep(1 * time.Second)
									if hp1 <= 0 {
										fmt.Println("You Died")
										check = 2
										break
									}	
								}else {
									fmt.Println("you took",attacken1,"damage")
									hp1 = hp1-attacken1
									time.Sleep(1 * time.Second)
									if hp1 <= 0 {
										fmt.Println("You Died")
										check = 2
										break
									}
								}
							} else {
								var b = time.Now().UnixNano()
								rand.Seed(b)
								var random3 = rand.Intn(5)
									if random3 == 1 || random3 == 2 || random3 == 3 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp1 = hp1 - attacken2 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}	
										}else {
											fmt.Println("you took",attacken2,"damage")
											hp1 = hp1-attacken2
											time.Sleep(1 * time.Second)	
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}
										}
									} else {
										fmt.Println("attack missed")
											time.Sleep(1 * time.Second)	
											continue
									}
								}
							}
							if check == 2 {
								time.Sleep(2 * time.Second)
								fmt.Println("Respawning")
								time.Sleep(5 * time.Second)
								continue
							}
						hp1 = hp1or
						attack1tr = 0
						hp1tr = 0
						attack2tr = 0
						attack3tr = 0
						attack4tr = 0
						montr = 0
						break
					}
						time.Sleep(4 * time.Second)
						fmt.Println("After beating the spartan, you continue up the mountain, and see a temple at the peak.")
						time.Sleep(4 * time.Second)
						fmt.Println("You realize that is where the guardian of the gem will be, and realize you might have to verse Zeus,")
						time.Sleep(4 * time.Second)
						fmt.Println("but you are comforted realizing Zeus is an ally.")
						time.Sleep(4 * time.Second)
						fmt.Println("With this assurance, you enter the temple.You are astonished to the sight of Zeus, but you notice he looks off.")
						time.Sleep(4 * time.Second)
						fmt.Println("He has succumbed to the Curse.")
						time.Sleep(4 * time.Second)
						for {
							hp1 = hp1 - hp1tr 
							hp1 = hp1or
							attack1 = attack1 - attack1tr
							mon = mon - montr
							hp1 = hp1 - hp1tr
							hp1or = hp1
							attack2 = attack2 - attack2tr
							attack3 = attack3 - attack3tr
							attack4 = attack4 - attack4tr
							hp2or = hp2
							attack1tr = 0
							hp1tr = 0
							attack2tr = 0
							attack3tr = 0
							attack4tr = 0
							montr = 0
							check = 1
							hp2 = 300
							hp2or = hp2
							attacken1 = 60
							attacken2 = 90
							attacken3 = 30
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
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 3 {
											hp2 = hp2 - attack4 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
											hp2 = hp2 - attack4
											fmt.Println("Enemy took",attack4,"damage")
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
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
								fmt.Println()
								fmt.Println("choose your attack")
								fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
								fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
								fmt.Println("3.heal",attack3,"health 100% accuracy")
								fmt.Scanln(&a)
								if a == 1 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp2 = hp2 - attack1 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									}else {
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack1,"damage")
										hp2 = hp2-attack1
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
									}
								} else if a == 2 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(3)
									if random == 1 || random == 2 || random == 3 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 3 {
											hp2 = hp2 - attack2 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
										time.Sleep(1 * time.Second)
										hp2 = hp2 - attack2
										fmt.Println("Enemy took",attack2,"damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
										}
									} else {
										time.Sleep(1 * time.Second)
										fmt.Println("attack missed")
										time.Sleep(1 * time.Second)
									}
								} else if a == 3 {
									if hp1 == hp1or {
										time.Sleep(1 * time.Second)
										fmt.Println("You're not low enough to heal")
										time.Sleep(1 * time.Second)
									} else if hp1 >= (hp1or - attack3){
										time.Sleep(1 * time.Second)
										hp1 = hp1or
										fmt.Println("You have been healed to max")
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
								time.Sleep(1 * time.Second)
								fmt.Println("enemy is attacking")
								time.Sleep(1 * time.Second)
								if hp2 <= (hp2or / 2) && hp1 > attacken2 {
									time.Sleep(1 * time.Second)
									hp2 = hp2 + attacken3
									fmt.Println("enemy has been healed",attacken3,"health")
									time.Sleep(1 * time.Second)
									continue
								}
								if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp1 = hp1 - attacken1 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											check = 2
											break
										}	
									}else {
										fmt.Println("you took",attacken1,"damage")
										hp1 = hp1-attacken1
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											check = 2
											break
										}
									}
								} else {
									var b = time.Now().UnixNano()
									rand.Seed(b)
									var random3 = rand.Intn(5)
										if random3 == 1 || random3 == 2 || random3 == 3 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp1 = hp1 - attacken2 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2
													break
												}	
											}else {
												fmt.Println("you took",attacken2,"damage")
												hp1 = hp1-attacken2
												time.Sleep(1 * time.Second)	
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2
													break
												}
											}
										} else {
											fmt.Println("attack missed")
												time.Sleep(1 * time.Second)	
												continue
										}
									}
								}
								if check == 2 {
									time.Sleep(2 * time.Second)
									fmt.Println("Respawning")
									time.Sleep(5 * time.Second)
									continue
								}
							hp1 = hp1or
							attack1tr = 0
							hp1tr = 0
							attack2tr = 0
							attack3tr = 0
							attack4tr = 0
							montr = 0
							break
						}
							attack3 = attack3 + 15
							time.Sleep(1 * time.Second)
							fmt.Println("You Don't kill Zeus, you have knocked the curse out of him. He recovers and thanks you.")
							time.Sleep(4 * time.Second)
							fmt.Println("He gives you some of his magic")
							time.Sleep(4 * time.Second)
							fmt.Println("Magic Upgraded")
							time.Sleep(4 * time.Second)
							fmt.Println("You take the gem, and proceed to the Cursed Castle.")
							time.Sleep(7 * time.Second)
							for space := 1 ; space <=100 ; space++{
								fmt.Println()
							}
							fmt.Println("MAP 3: Cursed Castle")
							for space := 1 ; space <=20 ; space++{
								fmt.Println()
							}
							time.Sleep(2 * time.Second)	
							for space := 1 ; space <=100 ; space++{
								fmt.Println()
							}
							time.Sleep(5 * time.Second)
							fmt.Println("The Castle Gate creaks open, as you take confident steps, entering the castle.")
							time.Sleep(5 * time.Second)
							fmt.Println("The castle is dark, and cursed. A Royal Servant stops you in your path, in the path leading to the main gates.")
							time.Sleep(5 * time.Second)
							fmt.Println("How dare you enter in this blissful castle. Your presence brings shame.")
							time.Sleep(4 * time.Second)
							for {
								hp1 = hp1 - hp1tr 
								hp1 = hp1or
								attack1 = attack1 - attack1tr
								mon = mon - montr
								hp1 = hp1 - hp1tr
								hp1or = hp1
								attack2 = attack2 - attack2tr
								attack3 = attack3 - attack3tr
								attack4 = attack4 - attack4tr
								hp2or = hp2
								attack1tr = 0
								hp1tr = 0
								attack2tr = 0
								attack3tr = 0
								attack4tr = 0
								montr = 0
								check = 1
								hp2 = 200
								hp2or = hp2
								attacken1 = 40
								attacken2 = 70
								attacken3 = 30
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
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack4 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
												hp2 = hp2 - attack4
												fmt.Println("Enemy took",attack4,"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
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
									fmt.Println()
									fmt.Println("choose your attack")
									fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
									fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
									fmt.Println("3.heal",attack3,"health 100% accuracy")
									fmt.Scanln(&a)
									if a == 1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp2 = hp2 - attack1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1,"damage")
											hp2 = hp2-attack1
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
										}
									} else if a == 2 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(3)
										if random == 1 || random == 2 || random == 3 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack2 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
											time.Sleep(1 * time.Second)
											hp2 = hp2 - attack2
											fmt.Println("Enemy took",attack2,"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
											}
										} else {
											time.Sleep(1 * time.Second)
											fmt.Println("attack missed")
											time.Sleep(1 * time.Second)
										}
									} else if a == 3 {
										if hp1 == hp1or {
											time.Sleep(1 * time.Second)
											fmt.Println("You're not low enough to heal")
											time.Sleep(1 * time.Second)
										} else if hp1 >= (hp1or - attack3){
											time.Sleep(1 * time.Second)
											hp1 = hp1or
											fmt.Println("You have been healed to max")
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
									time.Sleep(1 * time.Second)
									fmt.Println("enemy is attacking")
									time.Sleep(1 * time.Second)
									if hp2 <= (hp2or / 2) && hp1 > attacken2 {
										time.Sleep(1 * time.Second)
										hp2 = hp2 + attacken3
										fmt.Println("enemy has been healed",attacken3,"health")
										time.Sleep(1 * time.Second)
										continue
									}
									if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp1 = hp1 - attacken1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}	
										}else {
											fmt.Println("you took",attacken1,"damage")
											hp1 = hp1-attacken1
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}
										}
									} else {
										var b = time.Now().UnixNano()
										rand.Seed(b)
										var random3 = rand.Intn(5)
											if random3 == 1 || random3 == 2 || random3 == 3 {
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 0 {
													hp1 = hp1 - attacken2 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
													time.Sleep(1 * time.Second)
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2
														break
													}	
												}else {
													fmt.Println("you took",attacken2,"damage")
													hp1 = hp1-attacken2
													time.Sleep(1 * time.Second)	
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2
														break
													}
												}
											} else {
												fmt.Println("attack missed")
													time.Sleep(1 * time.Second)	
													continue
											}
										}
									}
									if check == 2 {
										time.Sleep(2 * time.Second)
										fmt.Println("Respawning")
										time.Sleep(5 * time.Second)
										continue
									}
								hp1 = hp1or
								hp1 = hp1or
								attack1tr = 0
								hp1tr = 0
								attack2tr = 0
								attack3tr = 0
								attack4tr = 0
								montr = 0
								break
							}
							mon = mon + 30
							fmt.Println("You gained 30 coins")
							time.Sleep(4 * time.Second)
							fmt.Println("Total coins:",mon)
							time.Sleep(4 * time.Second)
							fmt.Println("He yells Die Cursed Scum as he dies. He was wielding a steel chest plate.")
							time.Sleep(4 * time.Second)
							fmt.Println("You claim it")
							time.Sleep(4 * time.Second)
							fmt.Println("Armour upgraded")
							time.Sleep(4 * time.Second)
							fmt.Println("You carry on through the tight corridors, and encounter a sad weapon smith.")
							time.Sleep(5 * time.Second)
							fmt.Println("He is creating a sword, longer than you have ever seen. He has a pile of a variety of weapons. you approach him.")
							time.Sleep(5 * time.Second)
							fmt.Println("Greatings ashen one, He says while looking up from his anvil")
							time.Sleep(5 * time.Second)
							fmt.Println("He hands you a purple potion")
							time.Sleep(5 * time.Second)
							fmt.Println("You will need this for your journey ahead")
							time.Sleep(5 * time.Second)
							fmt.Println("You thank him and continue forwards")
							time.Sleep(5 * time.Second)
							var poip int
							var cursep int
							var hpst int
							var rpoip = 0
							var rcursep = 0
							poip = poip + 1
							fmt.Println("Poison potion collected")
							time.Sleep(5 * time.Second)
							fmt.Println("As you continue forwards you realize spread throughout the entire castle is a mercantile system.")
							time.Sleep(5 * time.Second)
							for {
								time.Sleep(2 * time.Second)
								fmt.Println("Which shop would you like to go?")
								fmt.Println("Total Coins:",mon)
								fmt.Println("1.Weapon Shop")
								fmt.Println("2.Armor Shop")
								fmt.Println("3.Magic shop")
								fmt.Println("4.Dark potions shop")
								fmt.Println("5.Exit")
								fmt.Scanln(&a)
								time.Sleep(2 * time.Second)
								if a == 1 {
									fmt.Println("What would you like to buy?")
									fmt.Println("Total Coins:",mon)
									fmt.Println("1.Light upgrade: 30 coins")
									fmt.Println("2.Heavy upgrade: 30 coins")
									fmt.Println("3.Ultimate upgrade: 30 coins")
									fmt.Println("4.Go back")
									fmt.Scanln(&b)
									time.Sleep(2 * time.Second)
									if b == 1 {
										if mon < 30 {
											time.Sleep(1 * time.Second)
											fmt.Println("You do not have enough to buy this item")
											time.Sleep(1 * time.Second)
										}else{
											mon = mon - 30
											fmt.Println("Light attack upgraded")
											attack1 = attack1 + 10
										}
									} else if b == 2 {
										if mon < 30 {
											time.Sleep(1 * time.Second)
											fmt.Println("You do not have enough to buy this item")
											time.Sleep(1 * time.Second)
										} else {
											mon = mon - 30
											fmt.Println("Heavy attack upgraded")
											attack2 = attack2 + 10
										}
									} else if b == 3 {
										if mon < 30 {
											time.Sleep(1 * time.Second)
											fmt.Println("You do not have enough to buy this item")
											time.Sleep(1 * time.Second)
										} else {
											mon = mon - 30
											fmt.Println("Ultimate attack upgraded")
											attack4 = attack4 + 15
										}
									} else {
										continue
									}
								} else if a == 2 {
									fmt.Println("What would you like to buy?")
									fmt.Println("Total Coins:",mon)
									fmt.Println("1.Helmet upgrade: 30 coins")
									fmt.Println("2.Chestplate upgrade: 40 coins")
									fmt.Println("3.Go back")
									fmt.Scanln(&b)
									time.Sleep(2 * time.Second)
									if b == 1 {
										if mon < 30 {
											time.Sleep(1 * time.Second)
											fmt.Println("You do not have enough to buy this item")
											time.Sleep(1 * time.Second)
										} else {
											mon = mon - 30
											fmt.Println("Helmet upgraded")
											hp1 = hp1 + 10
											hp1or = hp1
										}
									} else if b == 2 {
										if mon < 40 {
											time.Sleep(1 * time.Second)
											fmt.Println("You do not have enough to buy this item")
											time.Sleep(1 * time.Second)
										} else {
											mon = mon - 40
											fmt.Println("Chestplate upgraded")
											hp1 = hp1 + 20
											hp1or = hp1
										}
									} else {
										continue
									}									
								} else if a == 3 {
									fmt.Println("What would you like to buy?")
									fmt.Println("Total Coins:",mon)
									fmt.Println("1.Small magic upgrade: 30 coins")
									fmt.Println("2.Medium magic upgrade: 40 coins")
									fmt.Println("3.Large magic upgrade: 50 coins")
									fmt.Println("4.Go back")
									fmt.Scanln(&b)
									time.Sleep(2 * time.Second)
									if b == 1 {
										if mon < 30 {
											time.Sleep(1 * time.Second)
											fmt.Println("You do not have enough to buy this item")
											time.Sleep(1 * time.Second)
										} else {
											mon = mon - 30
											fmt.Println("Magic upgraded")
											attack3 = attack3 + 10
										}
									} else if b == 2 {
										if mon < 40 {
											time.Sleep(1 * time.Second)
											fmt.Println("You do not have enough to buy this item")
											time.Sleep(1 * time.Second)
										} else {
											mon = mon - 40
											fmt.Println("Magic upgraded")
											attack3 = attack3 + 15
										}
									} else if b == 3 {
										if mon < 50 {
											time.Sleep(1 * time.Second)
											fmt.Println("You do not have enough to buy this item")
											time.Sleep(1 * time.Second)
										} else {
											mon = mon - 50
											fmt.Println("Magic upgraded")
											attack3 = attack3 + 20
										}
									} else {
										continue
									}
								} else if a == 4{
									fmt.Println("What would you like to buy?")
									fmt.Println("Total Coins:",mon)
									fmt.Println("1.Poison potion: 20 coins")
									fmt.Println("2.Curse potion: 30 coins")
									fmt.Println("3.Life steal potion: 40 coins")
									fmt.Println("4.Go back")
									fmt.Scanln(&b)
									time.Sleep(2 * time.Second)
									if b == 1 {
										if mon < 20 {
											time.Sleep(1 * time.Second)
											fmt.Println("You do not have enough to buy this item")
											time.Sleep(1 * time.Second)
										} else {
											mon = mon - 20
											fmt.Println("Dark potion added")
											poip = poip + 1
										}
									} else if b == 2 {
										if mon < 30 {
											time.Sleep(1 * time.Second)
											fmt.Println("You do not have enough to buy this item")
											time.Sleep(1 * time.Second)
										} else {
											mon = mon - 30
											fmt.Println("Dark potion added")
											cursep = cursep + 1
										}
									} else if b == 3 {
										if mon < 40 {
											time.Sleep(1 * time.Second)
											fmt.Println("You do not have enough to buy this item")
											time.Sleep(1 * time.Second)
										} else {
											mon = mon - 40
											fmt.Println("Dark potion added")
											hpst = hpst + 1
										}
									} else {
										continue
									}
								} else {
									time.Sleep(2 * time.Second)
									break
								}
							}
							for {
								check = 1
								fmt.Println("You buy what you need, with the gold you obtained from all previous battles.")
								time.Sleep(5 * time.Second)
								fmt.Println("The people of what was once known as Toadstool castle are being oppressed by the fallen king, sent by Hell Lord Hades to manage the castle.")
								time.Sleep(5 * time.Second)
								fmt.Println("The old king was beheaded in the main square of the castle. You must avenge him.")
								time.Sleep(5 * time.Second)
								fmt.Println("You continue on through the courtyard, until you realize that the castle is completely deserted of enemies. Where do you go?")
								fmt.Println("1.Throne room")
								fmt.Println("2.Exit Castle")
								fmt.Println("3.Enter Dungeon")
								fmt.Println("4.Hide")
								fmt.Scanln(&b)
								if b == 1 {
									fmt.Println("The Throne room is completely trapped.")
									time.Sleep(5 * time.Second)
									fmt.Println("You turn around to attempt to open the door, but you see the shopkeeper from Map 1. At first thought.")
									time.Sleep(5 * time.Second)
									fmt.Println("You think he is cursed, but he looks just like before. He pulls out a fiery sword. Long Live Hades, He utters")
									time.Sleep(5 * time.Second)
									fmt.Println("He pushes you to the ground and points the sword towards your throat.")
									time.Sleep(5 * time.Second)
									fmt.Println("At this moment, you realize that Mystic Woods was a scapegoat, to lure you into a trap.")
									time.Sleep(5 * time.Second)
									fmt.Println("The shopkeeper stabs the flaming sword into your back, you feel a wave of pain until you fall into a forever darkness.")
									time.Sleep(3 * time.Second)
									check = 2
								} else if b == 2 {
									fmt.Println("You exit the castle, but enemies pour out of the bushes and kill you.")
									check = 2
								} else if b == 3 {
									fmt.Println("You enter the dungeons, and see the old king, he is in bad shape.")
									time.Sleep(5 * time.Second)
									fmt.Println("He speaks to you about the Fallen King's weakness, and you now feel confident in beating him.")
									time.Sleep(5 * time.Second)
									fmt.Println("Light Attacks. You thank the King, and vow to free him when the threat is over. You proceed up the stairs.")
									time.Sleep(3 * time.Second)
								} else {
									fmt.Println("You hide but you are found by guards and executed.")
									check = 2
								}
								if check == 2 {
									time.Sleep(2 * time.Second)
									fmt.Println("Respawning")
									time.Sleep(5 * time.Second)
									continue
								}
								break
							}
							time.Sleep(3 * time.Second)
							fmt.Println("After proceeding up the stairs, you enter the Fallen kings lair, and see a terrible dream. Betrayal.")
							time.Sleep(5 * time.Second)
							fmt.Println("You ignore it and carry on. You search to fight the fallen king, but encounter A Royal Servant!")
							for {
								hp1 = hp1 - hp1tr 
								hp1 = hp1or
								attack1 = attack1 - attack1tr
								mon = mon - montr
								hp1 = hp1 - hp1tr
								hp1or = hp1
								attack2 = attack2 - attack2tr
								attack3 = attack3 - attack3tr
								attack4 = attack4 - attack4tr
								hp2or = hp2
								attack1tr = 0
								hp1tr = 0
								attack2tr = 0
								attack3tr = 0
								attack4tr = 0
								montr = 0
								check = 1
								rpoip = 0
								rcursep = 0
								hp2 = 350
								hp2or = hp2
								attacken1 = 50
								attacken2 = 80
								attacken3 = 30
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
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack4 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
												hp2 = hp2 - attack4
												fmt.Println("Enemy took",attack4,"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
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
									fmt.Println()
									fmt.Println("choose your attack")
									fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
									fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
									fmt.Println("3.heal",attack3,"health 100% accuracy")
									fmt.Println("4.dark magic 100% accuracy")
									fmt.Scanln(&a)
									if a == 1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp2 = hp2 - attack1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1,"damage")
											hp2 = hp2-attack1
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
										}
									} else if a == 2 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(3)
										if random == 1 || random == 2 || random == 3 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack2 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
											time.Sleep(1 * time.Second)
											hp2 = hp2 - attack2
											fmt.Println("Enemy took",attack2,"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
											}
										} else {
											time.Sleep(1 * time.Second)
											fmt.Println("attack missed")
											time.Sleep(1 * time.Second)
										}
									} else if a == 3 {
										if hp1 == hp1or {
											time.Sleep(1 * time.Second)
											fmt.Println("You're not low enough to heal")
											time.Sleep(1 * time.Second)
										} else if hp1 >= (hp1or - attack3){
											time.Sleep(1 * time.Second)
											hp1 = hp1or
											fmt.Println("You have been healed to max")
											time.Sleep(1 * time.Second)
										} else {
											time.Sleep(1 * time.Second)
											hp1 = hp1 + attack3
											fmt.Println("You have been healed",attack3,"health")
											time.Sleep(1 * time.Second)
										}
									} else if a == 4 {
										time.Sleep(1 * time.Second)
										fmt.Println("Which potion would you like to use")
										fmt.Println("1.Poison:",poip)
										fmt.Println("2.Curse",cursep)
										fmt.Println("3.Life Steal",hpst)
										fmt.Scanln(&b)
										time.Sleep(1 * time.Second)
										if b == 1 {
											if poip == 0 {
												time.Sleep(1 * time.Second)
												fmt.Println("Sorry you don't have this item")
												time.Sleep(1 * time.Second)
												continue
											} else {
												poip--
												time.Sleep(1 * time.Second)
												rpoip = 5
											}
										} else if b == 2 {
											if cursep == 0 {
												time.Sleep(1 * time.Second)
												fmt.Println("Sorry you don't have this item")
												time.Sleep(1 * time.Second)
												continue
											} else {
												cursep--
												time.Sleep(1 * time.Second)
												rcursep = 10
											}
										} else if b == 3 {
											if hpst == 0 {
												time.Sleep(1 * time.Second)
												fmt.Println("Sorry you don't have this item")
												time.Sleep(1 * time.Second)
												continue
											} else {
												hpst--
												if hp1 == hp1or {
													hp2 = hp2 - 30
													time.Sleep(1 * time.Second)
													fmt.Println("You're not low enough to heal")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took 30 damage")
												} else if hp1 >= (hp1or - 30){
													time.Sleep(1 * time.Second)
													hp1 = hp1or
													hp2 = hp2 - 30
													fmt.Println("You have been healed to max")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took 30 damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}
												} else {
													time.Sleep(1 * time.Second)
													hp1 = hp1 + 30
													hp2 = hp2 - 30
													fmt.Println("You stole 30 hp")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}
												}
											}
										}
									} else {
										fmt.Println("Sorry that is an invalid command")
									}
									time.Sleep(1 * time.Second)
									fmt.Println("enemy is attacking")
									time.Sleep(1 * time.Second)
									if hp2 <= (hp2or / 2) && hp1 > attacken2 {
										time.Sleep(1 * time.Second)
										hp2 = hp2 + attacken3
										fmt.Println("enemy has been healed",attacken3,"health")
										time.Sleep(1 * time.Second)
										continue
									}
									if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp1 = hp1 - attacken1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2 
												break
											}	
										}else {
											fmt.Println("you took",attacken1,"damage")
											hp1 = hp1-attacken1
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2 
												break
											}
										}
									} else {
										var b = time.Now().UnixNano()
										rand.Seed(b)
										var random3 = rand.Intn(5)
											if random3 == 1 || random3 == 2 || random3 == 3 {
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 0 {
													hp1 = hp1 - attacken2 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
													time.Sleep(1 * time.Second)
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2 
														break
													}	
												}else {
													fmt.Println("you took",attacken2,"damage")
													hp1 = hp1-attacken2
													time.Sleep(1 * time.Second)	
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2 
														break
													}
												}
											} else {
												fmt.Println("attack missed")
													time.Sleep(1 * time.Second)	
													continue
											}
										}
										if rpoip != 0 {
											hp2 = hp2 - 10
											rpoip--
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy poisoned 10 damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}
										if rcursep != 0 {
											if rcursep == 1 {
												hp2 = hp2 - 150
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took 150 damage.")
												time.Sleep(1 * time.Second)
												rcursep = 0
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}
											rcursep--
											time.Sleep(1 * time.Second)
											fmt.Println(rcursep,"turns left")
											time.Sleep(1 * time.Second)
										}
									}
								hp1 = hp1or
								if check == 2 {
									time.Sleep(2 * time.Second)
									fmt.Println("Respawning")
									time.Sleep(5 * time.Second)
									continue
								}
							hp1 = hp1or
							hp1 = hp1or
							attack1tr = 0
							hp1tr = 0
							attack2tr = 0
							attack3tr = 0
							attack4tr = 0
							montr = 0
							break
							}
							time.Sleep(3 * time.Second)
							fmt.Println("After beating the servant, you realize he was different from the others.")
							time.Sleep(5 * time.Second)
							fmt.Println("He is one of the Fallen King's right-hand men.")
							time.Sleep(5 * time.Second)
							fmt.Println("He was leaving to kill the King in the dungeons.")
							time.Sleep(5 * time.Second)
							fmt.Println("Wondering where the Fallen King was, you exit the throne room, and see Terror, Death, and Blood.")
							time.Sleep(5 * time.Second)
							fmt.Println("The Merchants were killed, impaled on stakes for treachery. Enraged, you go back to the Throne Room.")
							time.Sleep(5 * time.Second)
							fmt.Println("The Fallen King is there and you get ready for battle")
							for {
								hp1 = hp1 - hp1tr 
								hp1 = hp1or
								attack1 = attack1 - attack1tr
								mon = mon - montr
								hp1 = hp1 - hp1tr
								hp1or = hp1
								attack2 = attack2 - attack2tr
								attack3 = attack3 - attack3tr
								attack4 = attack4 - attack4tr
								hp2or = hp2
								attack1tr = 0
								hp1tr = 0
								attack2tr = 0
								attack3tr = 0
								attack4tr = 0
								montr = 0
								check = 1
								rpoip = 0
								rcursep = 0
								hp2 = 1100
								hp2or = hp2
								attacken1 = 50
								attacken2 = 120
								attacken3 = 30
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
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack4 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
												hp2 = hp2 - attack4
												fmt.Println("Enemy took",attack4,"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
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
									fmt.Println()
									fmt.Println("choose your attack")
									fmt.Println("1.light attack ",attack1 * 2,"damage 100% accuracy")
									fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
									fmt.Println("3.heal",attack3,"health 100% accuracy")
									fmt.Println("4.dark magic 100% accuracy")
									fmt.Scanln(&a)
									if a == 1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp2 = hp2 - attack1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1 * 2 + (attack1),"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										} else {
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1 * 2,"damage")
											hp2 = hp2 - attack1 * 2
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
										}
									} else if a == 2 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(3)
										if random == 1 || random == 2 || random == 3 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack2 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
											time.Sleep(1 * time.Second)
											hp2 = hp2 - attack2
											fmt.Println("Enemy took",attack2,"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
											}
										} else {
											time.Sleep(1 * time.Second)
											fmt.Println("attack missed")
											time.Sleep(1 * time.Second)
										}
									} else if a == 3 {
										if hp1 == hp1or {
											time.Sleep(1 * time.Second)
											fmt.Println("You're not low enough to heal")
											time.Sleep(1 * time.Second)
										} else if hp1 >= (hp1or - attack3){
											time.Sleep(1 * time.Second)
											hp1 = hp1or
											fmt.Println("You have been healed to max")
											time.Sleep(1 * time.Second)
										} else {
											time.Sleep(1 * time.Second)
											hp1 = hp1 + attack3
											fmt.Println("You have been healed",attack3,"health")
											time.Sleep(1 * time.Second)
										}
									} else if a == 4 {
										time.Sleep(1 * time.Second)
										fmt.Println("Which potion would you like to use")
										fmt.Println("1.Poison:",poip)
										fmt.Println("2.Curse",cursep)
										fmt.Println("3.Life Steal",hpst)
										fmt.Scanln(&b)
										time.Sleep(1 * time.Second)
										if b == 1 {
											if poip == 0 {
												time.Sleep(1 * time.Second)
												fmt.Println("Sorry you don't have this item")
												time.Sleep(1 * time.Second)
												continue
											} else {
												poip--
												time.Sleep(1 * time.Second)
												rpoip = 5
											}
										} else if b == 2 {
											if cursep == 0 {
												time.Sleep(1 * time.Second)
												fmt.Println("Sorry you don't have this item")
												time.Sleep(1 * time.Second)
												continue
											} else {
												cursep--
												time.Sleep(1 * time.Second)
												rcursep = 10
											}
										} else if b == 3 {
											if hpst == 0 {
												time.Sleep(1 * time.Second)
												fmt.Println("Sorry you don't have this item")
												time.Sleep(1 * time.Second)
												continue
											} else {
												hpst--
												if hp1 == hp1or {
													hp2 = hp2 - 30
													time.Sleep(1 * time.Second)
													fmt.Println("You're not low enough to heal")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took 30 damage")
												} else if hp1 >= (hp1or - 30){
													time.Sleep(1 * time.Second)
													hp1 = hp1or
													hp2 = hp2 - 30
													fmt.Println("You have been healed to max")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took 30 damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}
												} else {
													time.Sleep(1 * time.Second)
													hp1 = hp1 + 30
													hp2 = hp2 - 30
													fmt.Println("You stole 30 hp")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}
												}
											}
										}
									} else {
										fmt.Println("Sorry that is an invalid command")
									}
									time.Sleep(1 * time.Second)
									fmt.Println("enemy is attacking")
									time.Sleep(1 * time.Second)
									if hp2 <= 300 && hp1 > attacken2 {
										time.Sleep(1 * time.Second)
										hp2 = hp2 + attacken3
										fmt.Println("enemy has been healed",attacken3,"health")
										time.Sleep(1 * time.Second)
										continue
									}
									if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp1 = hp1 - attacken1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2 
												break
											}	
										} else {
											fmt.Println("you took",attacken1,"damage")
											hp1 = hp1-attacken1
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2 
												break
											}
										}
									} else {
										var b = time.Now().UnixNano()
										rand.Seed(b)
										var random3 = rand.Intn(5)
											if random3 == 1 || random3 == 2 || random3 == 3 {
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 0 {
													hp1 = hp1 - attacken2 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
													time.Sleep(1 * time.Second)
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2 
														break
													}	
												}else {
													fmt.Println("you took",attacken2,"damage")
													hp1 = hp1-attacken2
													time.Sleep(1 * time.Second)	
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2 
														break
													}
												}
											} else {
												fmt.Println("attack missed")
													time.Sleep(1 * time.Second)	
													continue
											}
										}
										if rpoip != 0 {
											hp2 = hp2 - 10
											rpoip--
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy poisoned 10 damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}
										if rcursep != 0 {
											if rcursep == 1 {
												hp2 = hp2 - 150
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took 150 damage.")
												time.Sleep(1 * time.Second)
												rcursep = 0
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}
											rcursep--
											time.Sleep(1 * time.Second)
											fmt.Println(rcursep,"turns left")
											time.Sleep(1 * time.Second)
										}
									}
								hp1 = hp1or
								if check == 2 {
									time.Sleep(2 * time.Second)
									fmt.Println("Respawning")
									time.Sleep(5 * time.Second)
									continue
								}
							hp1 = hp1or
							hp1 = hp1or
							attack1tr = 0
							hp1tr = 0
							attack2tr = 0
							attack3tr = 0
							attack4tr = 0
							montr = 0
							break
							}
							fmt.Println("after the battle you turn around to see Hades. He effortlessly captures you alongside his men.")
							time.Sleep(5 * time.Second)
							for space := 1 ; space <=100 ; space++{
								fmt.Println()
							}
							fmt.Println("MAP 4: HELL'S GATE")
							for space := 1 ; space <=20 ; space++{
								fmt.Println()
							}
							time.Sleep(2 * time.Second)	
							for space := 1 ; space <=100 ; space++{
								fmt.Println()
							}
							time.Sleep(2 * time.Second)
							fmt.Println("You wake up the the endless bumping of the wheels of the caged cart rolling on the Lava terrain of Hell's Gate.")
							time.Sleep(5 * time.Second)
							fmt.Println("You are being rolled towards a castle, which is still about 2000m away. There is a sad-looking wolf in the cage.")
							time.Sleep(5 * time.Second)
							fmt.Println("Suddenly as you are about to doze off, the entire cage tips over. The wheel fell off.")
							time.Sleep(5 * time.Second)
							for { 
								fmt.Println("You see an opportunity to escape. Do you take it? ")
								fmt.Println("1.Yes")
								fmt.Println("2.No")
								fmt.Scanln(&b)
								time.Sleep(2 * time.Second)
								if b == 1 {
									fmt.Println("You force the lock open, and you and the wolf run, and the guards notice once you pick up speed.")
									time.Sleep(5 * time.Second)
									fmt.Println("The chase after you and shoot you with flame arrows, but they miss.")
									time.Sleep(5 * time.Second)
									fmt.Println("You run into Hellbent Woods, and locate a sleeping elite guard. Stealthily, you go around him and twist his neck.")
									time.Sleep(5 * time.Second)
									fmt.Println("You take his sword, and enter battle with the Guard.")
									attack2 = 130
									attack1 = 90
									attack4 = 200
								} else {
									time.Sleep(2 * time.Second)
									fmt.Println("The cart gets fixed, and Hades executes you ")
									time.Sleep(2 * time.Second)
									fmt.Println("Respawning")
									time.Sleep(5 * time.Second)
									continue
								}
								break
							}
							for{
								hp1 = hp1 - hp1tr 
								hp1 = hp1or
								attack1 = attack1 - attack1tr
								mon = mon - montr
								hp1 = hp1 - hp1tr
								hp1or = hp1
								attack2 = attack2 - attack2tr
								attack3 = attack3 - attack3tr
								attack4 = attack4 - attack4tr
								hp2or = hp2
								attack1tr = 0
								hp1tr = 0
								attack2tr = 0
								attack3tr = 0
								attack4tr = 0
								montr = 0
								check = 1
								rpoip = 0
								rcursep = 0
								hp2 = 400
								hp2or = hp2
								attacken1 = 60
								attacken2 = 100
								attacken3 = 40
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
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack4 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
												hp2 = hp2 - attack4
												fmt.Println("Enemy took",attack4,"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
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
									fmt.Println()
									fmt.Println("choose your attack")
									fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
									fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
									fmt.Println("3.heal",attack3,"health 100% accuracy")
									fmt.Println("4.dark magic 100% accuracy")
									fmt.Scanln(&a)
									if a == 1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp2 = hp2 - attack1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1,"damage")
											hp2 = hp2 - attack1
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
										}
									} else if a == 2 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(3)
										if random == 1 || random == 2 || random == 3 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack2 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
											time.Sleep(1 * time.Second)
											hp2 = hp2 - attack2
											fmt.Println("Enemy took",attack2,"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
											}
										} else {
											time.Sleep(1 * time.Second)
											fmt.Println("attack missed")
											time.Sleep(1 * time.Second)
										}
									} else if a == 3 {
										if hp1 == hp1or {
											time.Sleep(1 * time.Second)
											fmt.Println("You're not low enough to heal")
											time.Sleep(1 * time.Second)
										} else if hp1 >= (hp1or - attack3){
											time.Sleep(1 * time.Second)
											hp1 = hp1or
											fmt.Println("You have been healed to max")
											time.Sleep(1 * time.Second)
										} else {
											time.Sleep(1 * time.Second)
											hp1 = hp1 + attack3
											fmt.Println("You have been healed",attack3,"health")
											time.Sleep(1 * time.Second)
										}
									} else if a == 4 {
										time.Sleep(1 * time.Second)
										fmt.Println("Which potion would you like to use")
										fmt.Println("1.Poison:",poip)
										fmt.Println("2.Curse",cursep)
										fmt.Println("3.Life Steal",hpst)
										fmt.Scanln(&b)
										time.Sleep(1 * time.Second)
										if b == 1 {
											if poip == 0 {
												time.Sleep(1 * time.Second)
												fmt.Println("Sorry you don't have this item")
												time.Sleep(1 * time.Second)
												continue
											} else {
												poip--
												time.Sleep(1 * time.Second)
												rpoip = 5
											}
										} else if b == 2 {
											if cursep == 0 {
												time.Sleep(1 * time.Second)
												fmt.Println("Sorry you don't have this item")
												time.Sleep(1 * time.Second)
												continue
											} else {
												cursep--
												time.Sleep(1 * time.Second)
												rcursep = 10
											}
										} else if b == 3 {
											if hpst == 0 {
												time.Sleep(1 * time.Second)
												fmt.Println("Sorry you don't have this item")
												time.Sleep(1 * time.Second)
												continue
											} else {
												hpst--
												if hp1 == hp1or {
													hp2 = hp2 - 30
													time.Sleep(1 * time.Second)
													fmt.Println("You're not low enough to heal")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took 30 damage")
												} else if hp1 >= (hp1or - 30){
													time.Sleep(1 * time.Second)
													hp1 = hp1or
													hp2 = hp2 - 30
													fmt.Println("You have been healed to max")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took 30 damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}
												} else {
													time.Sleep(1 * time.Second)
													hp1 = hp1 + 30
													hp2 = hp2 - 30
													fmt.Println("You stole 30 hp")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}
												}
											}
										}
									} else {
										fmt.Println("Sorry that is an invalid command")
									}
									time.Sleep(1 * time.Second)
									fmt.Println("enemy is attacking")
									time.Sleep(1 * time.Second)
									if hp2 <= (hp2or / 2) && hp1 > attacken2 {
										time.Sleep(1 * time.Second)
										hp2 = hp2 + attacken3
										fmt.Println("enemy has been healed",attacken3,"health")
										time.Sleep(1 * time.Second)
										continue
									}
									if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp1 = hp1 - attacken1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2 
												break
											}	
										}else {
											fmt.Println("you took",attacken1,"damage")
											hp1 = hp1-attacken1
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2 
												break
											}
										}
									} else {
										var b = time.Now().UnixNano()
										rand.Seed(b)
										var random3 = rand.Intn(5)
											if random3 == 1 || random3 == 2 || random3 == 3 {
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 0 {
													hp1 = hp1 - attacken2 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
													time.Sleep(1 * time.Second)
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2 
														break
													}	
												}else {
													fmt.Println("you took",attacken2,"damage")
													hp1 = hp1-attacken2
													time.Sleep(1 * time.Second)	
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2 
														break
													}
												}
											} else {
												fmt.Println("attack missed")
													time.Sleep(1 * time.Second)	
													continue
											}
										}
										if rpoip != 0 {
											hp2 = hp2 - 10
											rpoip--
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy poisoned 10 damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}
										if rcursep != 0 {
											if rcursep == 1 {
												hp2 = hp2 - 150
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took 150 damage.")
												time.Sleep(1 * time.Second)
												rcursep = 0
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}
											rcursep--
											time.Sleep(1 * time.Second)
											fmt.Println(rcursep,"turns left")
											time.Sleep(1 * time.Second)
										}
									}
								if check == 2 {
									time.Sleep(2 * time.Second)
									fmt.Println("Respawning")
									time.Sleep(5 * time.Second)
									continue
								}
								if check == 2 {
									time.Sleep(2 * time.Second)
									fmt.Println("Respawning")
									time.Sleep(5 * time.Second)
									continue
								}
							hp1 = hp1or
							attack1tr = 0
							hp1tr = 0
							attack2tr = 0
							attack3tr = 0
							attack4tr = 0
							montr = 0
							break
							}
						    var petat = 50
							var pethe = 40
							var toadkiller = 0
							time.Sleep(2 * time.Second)
							fmt.Println("You continue with the wolf. The Wolf has now became your companion.")
							time.Sleep(5 * time.Second)
							fmt.Println("You reclaim your armor from the guard, and proceed towards the castle, where Hades is at the moment.")
							time.Sleep(5 * time.Second)
							fmt.Println("He expects you as a prisoner in his presence, but little does he know, he will see fate in a matter of time.")
							time.Sleep(5 * time.Second)
							fmt.Println("You run towards the castle, but feel something grab your arm. It is a Hades Council Member.")
							time.Sleep(5 * time.Second)
							fmt.Println("He says Greetings, Ashen One You realize he is a spy.")
							time.Sleep(5 * time.Second)
							fmt.Println("he takes you into the Woods, and see a system of refugees in a camp. There are shops there.")
							time.Sleep(5 * time.Second)
							var shop1 = 0
							var shop2 = 0
							var shop3 = 0
							var shop4 = 0
							var shop5 = 0
							for {
								fmt.Println("Where would you like to go?")
								fmt.Println("1.Pet shop")
								fmt.Println("2.Weapon shop")
								fmt.Println("3.Shady shop")
								fmt.Println("4.Armor shop")
								fmt.Println("5.Dark potion shop")
								fmt.Println("6.Exit")
								fmt.Scanln(&b)
								time.Sleep(1 * time.Second)
								if b == 1 {
									if shop1 == 0 {
										shop1++
										petat = petat + 10
										pethe = pethe + 10
										fmt.Println("The merchent gives your wolf some food and gives him some armor")
										time.Sleep(1 * time.Second)
									} else {
										fmt.Println("The merchent gives your wolf some food")
										time.Sleep(1 * time.Second)
									}
								} else if b == 2 {
									if shop2 == 0 {
										shop2++
										attack1 = attack1 + 10
										attack2 = attack2 + 10
										attack4 = attack4 + 20
										fmt.Println("The merchent sharpen your weapon blade")
										time.Sleep(1 * time.Second)
									} else {
										fmt.Println("The merchent inspects your weapon but it is still in good condition")
										time.Sleep(1 * time.Second)
									}
								} else if b == 3 {
									if shop3 == 0 {
										shop3++
										toadkiller = 1
										fmt.Println("The merchent hands you a strange potion Use this to beat the one with the shroom. He says")
										time.Sleep(5 * time.Second)
										fmt.Println("You look at the potion. It doesn't look like anything you've seen before. You thank the merchnt though and continue shopping.")
										time.Sleep(1 * time.Second)
										} else {
										fmt.Println("The merchent wishes you luck for the future")
										time.Sleep(1 * time.Second)
									}
								} else if b == 4 {
									if shop4 == 0 {
										shop4++
										hp1 = hp1 + 50
										hp1or = hp1
										fmt.Println("The merchent fixes your armor")
										time.Sleep(1 * time.Second)
									} else {
										fmt.Println("The merchent says hi")
										time.Sleep(1 * time.Second)
									}
								} else if b == 5 {
									if shop5 == 0 {
										shop5++
										cursep = cursep + 1
										fmt.Println("The merchent gives you a curse potion")
										time.Sleep(1 * time.Second)
									} else {
										fmt.Println("The merchent says hi")
										time.Sleep(1 * time.Second)
									}
								} else {
									time.Sleep(1 * time.Second)
									break
								}
							}
							if toadkiller == 0 {
								toadkiller = 1
								fmt.Println("As you exit a shop a merchent stops you. The merchent hands you a strange potion Use this to beat the one with the shroom. He says")
								time.Sleep(5 * time.Second)
								fmt.Println("You look at the potion. It doesn't look like anything you've seen before. You thank the merchnt though and continue shopping.")
								time.Sleep(1 * time.Second)
							}
							fmt.Println("After your business, you exit the Black Market, and are greeted by the heat of the pouring lava from the Magma.")
							time.Sleep(5 * time.Second)
							fmt.Println("You notice Hades castle, and head towards it...The castle gates open, and you are greeted by Hades' best man.")
							time.Sleep(5 * time.Second)
							fmt.Println("Halt, Who goes there! He shouts.")
							for {
								hp1 = hp1 - hp1tr 
								hp1 = hp1or
								attack1 = attack1 - attack1tr
								mon = mon - montr
								hp1 = hp1 - hp1tr
								hp1or = hp1
								attack2 = attack2 - attack2tr
								attack3 = attack3 - attack3tr
								attack4 = attack4 - attack4tr
								hp2or = hp2
								attack1tr = 0
								hp1tr = 0
								attack2tr = 0
								attack3tr = 0
								attack4tr = 0
								montr = 0
								check = 1
								rpoip = 0
								rcursep = 0
								hp2 = 600
								hp2or = hp2
								attacken1 = 70
								attacken2 = 110
								attacken3 = 40
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
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack4 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
												hp2 = hp2 - attack4
												fmt.Println("Enemy took",attack4,"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
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
									fmt.Println()
									fmt.Println("choose your attack")
									fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
									fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
									fmt.Println("3.heal",attack3,"health 100% accuracy")
									fmt.Println("4.dark magic 100% accuracy")
									fmt.Scanln(&a)
									if a == 1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp2 = hp2 - attack1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1,"damage")
											hp2 = hp2 - attack1
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
										}
									} else if a == 2 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(3)
										if random == 1 || random == 2 || random == 3 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack2 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
											time.Sleep(1 * time.Second)
											hp2 = hp2 - attack2
											fmt.Println("Enemy took",attack2,"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
											}
										} else {
											time.Sleep(1 * time.Second)
											fmt.Println("attack missed")
											time.Sleep(1 * time.Second)
										}
									} else if a == 3 {
										if hp1 == hp1or {
											time.Sleep(1 * time.Second)
											fmt.Println("You're not low enough to heal")
											time.Sleep(1 * time.Second)
										} else if hp1 >= (hp1or - attack3){
											time.Sleep(1 * time.Second)
											hp1 = hp1or
											fmt.Println("You have been healed to max")
											time.Sleep(1 * time.Second)
										} else {
											time.Sleep(1 * time.Second)
											hp1 = hp1 + attack3
											fmt.Println("You have been healed",attack3,"health")
											time.Sleep(1 * time.Second)
										}
									} else if a == 4 {
										time.Sleep(1 * time.Second)
										fmt.Println("Which potion would you like to use")
										fmt.Println("1.Poison:",poip)
										fmt.Println("2.Curse",cursep)
										fmt.Println("3.Life Steal",hpst)
										fmt.Scanln(&b)
										time.Sleep(1 * time.Second)
										if b == 1 {
											if poip == 0 {
												time.Sleep(1 * time.Second)
												fmt.Println("Sorry you don't have this item")
												time.Sleep(1 * time.Second)
												continue
											} else {
												poip--
												time.Sleep(1 * time.Second)
												rpoip = 5
											}
										} else if b == 2 {
											if cursep == 0 {
												time.Sleep(1 * time.Second)
												fmt.Println("Sorry you don't have this item")
												time.Sleep(1 * time.Second)
												continue
											} else {
												cursep--
												time.Sleep(1 * time.Second)
												rcursep = 10
											}
										} else if b == 3 {
											if hpst == 0 {
												time.Sleep(1 * time.Second)
												fmt.Println("Sorry you don't have this item")
												time.Sleep(1 * time.Second)
												continue
											} else {
												hpst--
												if hp1 == hp1or {
													hp2 = hp2 - 30
													time.Sleep(1 * time.Second)
													fmt.Println("You're not low enough to heal")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took 30 damage")
												} else if hp1 >= (hp1or - 30){
													time.Sleep(1 * time.Second)
													hp1 = hp1or
													hp2 = hp2 - 30
													fmt.Println("You have been healed to max")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took 30 damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}
												} else {
													time.Sleep(1 * time.Second)
													hp1 = hp1 + 30
													hp2 = hp2 - 30
													fmt.Println("You stole 30 hp")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}
												}
											}
										}
									} else {
										fmt.Println("Sorry that is an invalid command")
									}
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(5)
									if random == 3 {
										if hp1 == hp1or / 2 {
											hp1 = hp1 + pethe
											time.Sleep(1 * time.Second)
											fmt.Println("Your pet healed you",pethe,"hp")
										} else {
											hp2 = hp2 - petat
											time.Sleep(1 * time.Second)
											fmt.Println("Your pet did",petat,"damage")
										}
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									}
									time.Sleep(1 * time.Second)
									fmt.Println("enemy is attacking")
									time.Sleep(1 * time.Second)
									if hp2 <= (hp2or / 2) && hp1 > attacken2 {
										time.Sleep(1 * time.Second)
										hp2 = hp2 + attacken3
										fmt.Println("enemy has been healed",attacken3,"health")
										time.Sleep(1 * time.Second)
										continue
									}
									if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp1 = hp1 - attacken1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2 
												break
											}	
										}else {
											fmt.Println("you took",attacken1,"damage")
											hp1 = hp1-attacken1
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2 
												break
											}
										}
									} else {
										var b = time.Now().UnixNano()
										rand.Seed(b)
										var random3 = rand.Intn(5)
											if random3 == 1 || random3 == 2 || random3 == 3 {
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 0 {
													hp1 = hp1 - attacken2 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
													time.Sleep(1 * time.Second)
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2 
														break
													}	
												}else {
													fmt.Println("you took",attacken2,"damage")
													hp1 = hp1-attacken2
													time.Sleep(1 * time.Second)	
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2 
														break
													}
												}
											} else {
												fmt.Println("attack missed")
													time.Sleep(1 * time.Second)	
													continue
											}
										}
										if rpoip != 0 {
											hp2 = hp2 - 10
											rpoip--
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy poisoned 10 damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}
										if rcursep != 0 {
											if rcursep == 1 {
												hp2 = hp2 - 150
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took 150 damage.")
												time.Sleep(1 * time.Second)
												rcursep = 0
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}
											rcursep--
											time.Sleep(1 * time.Second)
											fmt.Println(rcursep,"turns left")
											time.Sleep(1 * time.Second)
										}
									}
									if check == 2 {
										time.Sleep(2 * time.Second)
										fmt.Println("Respawning")
										time.Sleep(5 * time.Second)
										continue
									}
								hp1 = hp1or
								attack1tr = 0
								hp1tr = 0
								attack2tr = 0
								attack3tr = 0
								attack4tr = 0
								montr = 0
								break
							}
							fmt.Println("After beating the Top Soldier, you claim all his armor.")
							time.Sleep(5 * time.Second)
							fmt.Println("You can go forge your armor with your existing armor, or continue. What do you do?")
							fmt.Println("1.Forge")
							fmt.Println("2.Don't forge")
							fmt.Scanln(&b)
							if b == 1 {
								hp1 = hp1 + 60
								hp1or = hp1
								time.Sleep(1 * time.Second)
								fmt.Println("You re-enter the black market, and ask the Blacksmith for the forging. He agrees, and wishes you luck.")
								time.Sleep(1 * time.Second)
							}
							time.Sleep(1 * time.Second)
							fmt.Println("Now, you are confident, so you go back to the Castle, and head straight for Hades throne.")
							time.Sleep(5 * time.Second)
							fmt.Println("Hades is sitting there, and is surprised to see you. How did you escape!")
							time.Sleep(5 * time.Second)
							fmt.Println("You say, Its a Work of art, something a horned creep wouldn't know")
							for {
								hp1 = hp1 - hp1tr 
								hp1 = hp1or
								attack1 = attack1 - attack1tr
								mon = mon - montr
								hp1 = hp1 - hp1tr
								hp1or = hp1
								attack2 = attack2 - attack2tr
								attack3 = attack3 - attack3tr
								attack4 = attack4 - attack4tr
								hp2or = hp2
								attack1tr = 0
								hp1tr = 0
								attack2tr = 0
								attack3tr = 0
								attack4tr = 0
								montr = 0
								check = 1
								rpoip = 0
								rcursep = 0
								hp2 = 750
								hp2or = hp2
								attacken1 = 80
								attacken2 = 120
								attacken3 = 50
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
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack4 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
												hp2 = hp2 - attack4
												fmt.Println("Enemy took",attack4,"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
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
									fmt.Println()
									fmt.Println("choose your attack")
									fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
									fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
									fmt.Println("3.heal",attack3,"health 100% accuracy")
									fmt.Println("4.dark magic 100% accuracy")
									fmt.Scanln(&a)
									if a == 1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp2 = hp2 - attack1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1,"damage")
											hp2 = hp2 - attack1
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
										}
									} else if a == 2 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(3)
										if random == 1 || random == 2 || random == 3 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack2 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
											time.Sleep(1 * time.Second)
											hp2 = hp2 - attack2
											fmt.Println("Enemy took",attack2,"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
											}
										} else {
											time.Sleep(1 * time.Second)
											fmt.Println("attack missed")
											time.Sleep(1 * time.Second)
										}
									} else if a == 3 {
										if hp1 == hp1or {
											time.Sleep(1 * time.Second)
											fmt.Println("You're not low enough to heal")
											time.Sleep(1 * time.Second)
										} else if hp1 >= (hp1or - attack3){
											time.Sleep(1 * time.Second)
											hp1 = hp1or
											fmt.Println("You have been healed to max")
											time.Sleep(1 * time.Second)
										} else {
											time.Sleep(1 * time.Second)
											hp1 = hp1 + attack3
											fmt.Println("You have been healed",attack3,"health")
											time.Sleep(1 * time.Second)
										}
									} else if a == 4 {
										time.Sleep(1 * time.Second)
										fmt.Println("Which potion would you like to use")
										fmt.Println("1.Poison:",poip)
										fmt.Println("2.Curse",cursep)
										fmt.Println("3.Life Steal",hpst)
										fmt.Scanln(&b)
										time.Sleep(1 * time.Second)
										if b == 1 {
											if poip == 0 {
												time.Sleep(1 * time.Second)
												fmt.Println("Sorry you don't have this item")
												time.Sleep(1 * time.Second)
												continue
											} else {
												poip--
												time.Sleep(1 * time.Second)
												rpoip = 5
											}
										} else if b == 2 {
											if cursep == 0 {
												time.Sleep(1 * time.Second)
												fmt.Println("Sorry you don't have this item")
												time.Sleep(1 * time.Second)
												continue
											} else {
												cursep--
												time.Sleep(1 * time.Second)
												rcursep = 10
											}
										} else if b == 3 {
											if hpst == 0 {
												time.Sleep(1 * time.Second)
												fmt.Println("Sorry you don't have this item")
												time.Sleep(1 * time.Second)
												continue
											} else {
												hpst--
												if hp1 == hp1or {
													hp2 = hp2 - 30
													time.Sleep(1 * time.Second)
													fmt.Println("You're not low enough to heal")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took 30 damage")
												} else if hp1 >= (hp1or - 30){
													time.Sleep(1 * time.Second)
													hp1 = hp1or
													hp2 = hp2 - 30
													fmt.Println("You have been healed to max")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took 30 damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}
												} else {
													time.Sleep(1 * time.Second)
													hp1 = hp1 + 30
													hp2 = hp2 - 30
													fmt.Println("You stole 30 hp")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}
												}
											}
										}
									} else {
										fmt.Println("Sorry that is an invalid command")
									}
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(5)
									if random == 3 {
										if hp1 == hp1or / 2 {
											hp1 = hp1 + pethe
											time.Sleep(1 * time.Second)
											fmt.Println("Your pet healed you",pethe,"hp")
										} else {
											hp2 = hp2 - petat
											time.Sleep(1 * time.Second)
											fmt.Println("Your pet did",petat,"damage")
										}
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									}
									time.Sleep(1 * time.Second)
									fmt.Println("enemy is attacking")
									time.Sleep(1 * time.Second)
									if hp2 <= (hp2or / 3) && hp1 > attacken2 {
										time.Sleep(1 * time.Second)
										hp2 = hp2 + attacken3
										fmt.Println("enemy has been healed",attacken3,"health")
										time.Sleep(1 * time.Second)
										continue
									}
									if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp1 = hp1 - attacken1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2 
												break
											}	
										}else {
											fmt.Println("you took",attacken1,"damage")
											hp1 = hp1-attacken1
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2 
												break
											}
										}
									} else {
										var b = time.Now().UnixNano()
										rand.Seed(b)
										var random3 = rand.Intn(5)
											if random3 == 1 || random3 == 2 || random3 == 3 {
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 0 {
													hp1 = hp1 - attacken2 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
													time.Sleep(1 * time.Second)
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2 
														break
													}	
												}else {
													fmt.Println("you took",attacken2,"damage")
													hp1 = hp1-attacken2
													time.Sleep(1 * time.Second)	
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2 
														break
													}
												}
											} else {
												fmt.Println("attack missed")
													time.Sleep(1 * time.Second)	
													continue
											}
										}
										if rpoip != 0 {
											hp2 = hp2 - 10
											rpoip--
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy poisoned 10 damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}
										if rcursep != 0 {
											if rcursep == 1 {
												hp2 = hp2 - 150
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took 150 damage.")
												time.Sleep(1 * time.Second)
												rcursep = 0
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}
											rcursep--
											time.Sleep(1 * time.Second)
											fmt.Println(rcursep,"turns left")
											time.Sleep(1 * time.Second)
										}
									}
									if check == 2 {
										time.Sleep(2 * time.Second)
										fmt.Println("Respawning")
										time.Sleep(5 * time.Second)
										continue
									}
								hp1 = hp1or
								attack1tr = 0
								hp1tr = 0
								attack2tr = 0
								attack3tr = 0
								attack4tr = 0
								montr = 0
								break
							}
							time.Sleep(1 * time.Second)
							fmt.Println("After Beating Hades, He drops down.")
							time.Sleep(5 * time.Second)
							fmt.Println("HOW!, he screams as he dies, but something is wrong, the curse hasn't worn off.")
							time.Sleep(5 * time.Second)
							fmt.Println("Right as you realize, the Whole castle comes crashing down and you fall over into darkness.")
							time.Sleep(5 * time.Second)
							for space := 1 ; space <=100 ; space++{
								fmt.Println()
							}
							fmt.Println("MAP 5 : TOADSTOOL VILLAGE")
							for space := 1 ; space <=20 ; space++{
								fmt.Println()
							}
							time.Sleep(2 * time.Second)	
							for space := 1 ; space <=100 ; space++{
								fmt.Println()
							}
							for {
								fmt.Println("You wake up, lost, in a forest, with no memory to recall.")
								time.Sleep(4 * time.Second)
								fmt.Println("The leaves of the trees, blocked the sunlight from the two suns.")
								time.Sleep(4 * time.Second)
								fmt.Println("You notice a sword in your hand, in horrible condition.") 
								time.Sleep(4 * time.Second)
								fmt.Println("Confused, you get up, despite the physical pain, You march forward in hope to find life, but it is hopeless.")
								time.Sleep(4 * time.Second)
								fmt.Println("Suddenly, you are shocked with an impact to the back of your head. It's a toad.")
								time.Sleep(4 * time.Second)
								fmt.Println("HAHAHAHAHHA! Laughs a sinister voice. My minion almost KILLED You! I'm TOAD MAN!")
								time.Sleep(4 * time.Second)
								fmt.Println("Toad man attacks right.")
								fmt.Println("1.Block left")
								fmt.Println("2.Block right")
								fmt.Println("3.Duck")
								fmt.Scanln(&b)
								time.Sleep(2 * time.Second)
									check = 1
									if b == 1 {
										fmt.Println("Toad man killed you")
										time.Sleep(2 * time.Second)
										fmt.Println("Respawing")
										time.Sleep(4 * time.Second)
										continue
									} else if b == 2 {
										fmt.Println("You blocked Toad man's attack")
									} else {
										fmt.Println("Toad man killed you")
										time.Sleep(2 * time.Second)
										fmt.Println("Respawing")
										time.Sleep(4 * time.Second)
										continue
									}
									time.Sleep(2 * time.Second)	
									fmt.Println("YOU WONT BEAT ME!")
									time.Sleep(2 * time.Second)	
									fmt.Println("Toad man uses lightning bolt.")
									fmt.Println("1.Block left")
									fmt.Println("2.Block right")
									fmt.Println("3.Roll")
									fmt.Scanln(&b)
									time.Sleep(2 * time.Second)	
									if b == 1 {
										fmt.Println("Toad man killed you")
										time.Sleep(2 * time.Second)
										fmt.Println("Respawing")
										time.Sleep(4 * time.Second)
										continue
									} else if b == 2 {
										fmt.Println("Toad man killed you")
										time.Sleep(2 * time.Second)
										fmt.Println("Respawing")
										time.Sleep(4 * time.Second)
										continue
									} else {
										fmt.Println("You rolled under the lightning bolt")
									}
									time.Sleep(2 * time.Second)	
									fmt.Println("EVIL SHALL RULE ALL!")
									time.Sleep(2 * time.Second)	
									fmt.Println("Toad Man Slashes Through middle!")
									fmt.Println("1.Parry")
									fmt.Println("2.Duck")
									fmt.Println("3.Block left")
									fmt.Scanln(&b)
									time.Sleep(2 * time.Second)	
									if b == 1 {
										fmt.Println("You dodged his attack")
									} else if b == 2 {
										fmt.Println("Toad man killed you")
										time.Sleep(2 * time.Second)
										fmt.Println("Respawing")
										time.Sleep(4 * time.Second)
										continue
									} else {
										fmt.Println("Toad man killed you")
										time.Sleep(2 * time.Second)
										fmt.Println("Respawing")
										time.Sleep(4 * time.Second)
										continue				
									}
									time.Sleep(2 * time.Second)	
									fmt.Println("Toad man uses and ariel attack. DIE!")
									fmt.Println("1.Uppercut")
									fmt.Println("2.Dive")
									fmt.Println("3.Parry")
									fmt.Scanln(&b)
									time.Sleep(2 * time.Second)	
									if b == 1 {
										fmt.Println("You punch Toad man in the face.")
									} else if b == 2 {
										fmt.Println("Toad man killed you")
										time.Sleep(2 * time.Second)
										fmt.Println("Respawing")
										time.Sleep(4 * time.Second)
										continue
									} else {
										fmt.Println("Toad man killed you")
										time.Sleep(2 * time.Second)
										fmt.Println("Respawing")
										time.Sleep(4 * time.Second)
										continue
									}
									break
							}
							time.Sleep(2 * time.Second)
							fmt.Println("he is quite hurt from all your counters, and he takes one last stand.")
							for {
								hp1 = hp1 - hp1tr 
								hp1 = hp1or
								attack1 = attack1 - attack1tr
								mon = mon - montr
								hp1 = hp1 - hp1tr
								hp1or = hp1
								attack2 = attack2 - attack2tr
								attack3 = attack3 - attack3tr
								attack4 = attack4 - attack4tr
								hp2or = hp2
								attack1tr = 0
								hp1tr = 0
								attack2tr = 0
								attack3tr = 0
								attack4tr = 0
								montr = 0
								check = 1
								rpoip = 0
								rcursep = 0
								hp2 = 800
								hp2or = hp2
								attacken1 = 80
								attacken2 = 120
								attacken3 = 50
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
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack4 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
												hp2 = hp2 - attack4
												fmt.Println("Enemy took",attack4,"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
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
									fmt.Println()
									fmt.Println("choose your attack")
									fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
									fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
									fmt.Println("3.heal",attack3,"health 100% accuracy")
									fmt.Println("4.dark magic 100% accuracy")
									fmt.Scanln(&a)
									if a == 1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp2 = hp2 - attack1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1,"damage")
											hp2 = hp2 - attack1
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
										}
									} else if a == 2 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(3)
										if random == 1 || random == 2 || random == 3 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack2 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
											time.Sleep(1 * time.Second)
											hp2 = hp2 - attack2
											fmt.Println("Enemy took",attack2,"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
											}
										} else {
											time.Sleep(1 * time.Second)
											fmt.Println("attack missed")
											time.Sleep(1 * time.Second)
										}
									} else if a == 3 {
										if hp1 == hp1or {
											time.Sleep(1 * time.Second)
											fmt.Println("You're not low enough to heal")
											time.Sleep(1 * time.Second)
										} else if hp1 >= (hp1or - attack3){
											time.Sleep(1 * time.Second)
											hp1 = hp1or
											fmt.Println("You have been healed to max")
											time.Sleep(1 * time.Second)
										} else {
											time.Sleep(1 * time.Second)
											hp1 = hp1 + attack3
											fmt.Println("You have been healed",attack3,"health")
											time.Sleep(1 * time.Second)
										}
									} else if a == 4 {
										time.Sleep(1 * time.Second)
										fmt.Println("Which potion would you like to use")
										fmt.Println("1.Poison:",poip)
										fmt.Println("2.Curse",cursep)
										fmt.Println("3.Life Steal",hpst)
										fmt.Scanln(&b)
										time.Sleep(1 * time.Second)
										if b == 1 {
											if poip == 0 {
												time.Sleep(1 * time.Second)
												fmt.Println("Sorry you don't have this item")
												time.Sleep(1 * time.Second)
												continue
											} else {
												poip--
												time.Sleep(1 * time.Second)
												rpoip = 5
											}
										} else if b == 2 {
											if cursep == 0 {
												time.Sleep(1 * time.Second)
												fmt.Println("Sorry you don't have this item")
												time.Sleep(1 * time.Second)
												continue
											} else {
												cursep--
												time.Sleep(1 * time.Second)
												rcursep = 10
											}
										} else if b == 3 {
											if hpst == 0 {
												time.Sleep(1 * time.Second)
												fmt.Println("Sorry you don't have this item")
												time.Sleep(1 * time.Second)
												continue
											} else {
												hpst--
												if hp1 == hp1or {
													hp2 = hp2 - 30
													time.Sleep(1 * time.Second)
													fmt.Println("You're not low enough to heal")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took 30 damage")
												} else if hp1 >= (hp1or - 30){
													time.Sleep(1 * time.Second)
													hp1 = hp1or
													hp2 = hp2 - 30
													fmt.Println("You have been healed to max")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took 30 damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}
												} else {
													time.Sleep(1 * time.Second)
													hp1 = hp1 + 30
													hp2 = hp2 - 30
													fmt.Println("You stole 30 hp")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}
												}
											}
										}
									} else {
										fmt.Println("Sorry that is an invalid command")
									}
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(5)
									if random == 3 {
										if hp1 == hp1or / 2 {
											hp1 = hp1 + pethe
											time.Sleep(1 * time.Second)
											fmt.Println("You have been healed",pethe,"hp")
										} else {
											hp2 = hp2 - petat
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",petat,"damage")
										}
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									}
									time.Sleep(1 * time.Second)
									fmt.Println("enemy is attacking")
									time.Sleep(1 * time.Second)
									if hp2 <= (hp2or / 3) && hp1 > attacken2 {
										time.Sleep(1 * time.Second)
										hp2 = hp2 + attacken3
										fmt.Println("enemy has been healed",attacken3,"health")
										time.Sleep(1 * time.Second)
										continue
									}
									if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp1 = hp1 - attacken1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}	
										}else {
											fmt.Println("you took",attacken1,"damage")
											hp1 = hp1-attacken1
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}
										}
									} else {
										var b = time.Now().UnixNano()
										rand.Seed(b)
										var random3 = rand.Intn(5)
											if random3 == 1 || random3 == 2 || random3 == 3 {
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 0 {
													hp1 = hp1 - attacken2 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
													time.Sleep(1 * time.Second)
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2
														break
													}	
												}else {
													fmt.Println("you took",attacken2,"damage")
													hp1 = hp1-attacken2
													time.Sleep(1 * time.Second)	
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2
														break
													}
												}
											} else {
												fmt.Println("attack missed")
													time.Sleep(1 * time.Second)	
													continue
											}
										}
										if rpoip != 0 {
											hp2 = hp2 - 10
											rpoip--
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy poisoned 10 damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}
										if rcursep != 0 {
											if rcursep == 1 {
												hp2 = hp2 - 150
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took 150 damage.")
												time.Sleep(1 * time.Second)
												rcursep = 0
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}
											rcursep--
											time.Sleep(1 * time.Second)
											fmt.Println(rcursep,"turns left")
											time.Sleep(1 * time.Second)
										}
									}
									if check == 2 {
										time.Sleep(2 * time.Second)
										fmt.Println("Respawning")
										time.Sleep(5 * time.Second)
										continue
									}
									hp1 = hp1or
									attack1tr = 0
									hp1tr = 0
									attack2tr = 0
									attack3tr = 0
									attack4tr = 0
									montr = 0
									break
							}
							fmt.Println("Toad man is weakened and falls to the ground. You place your sword at him but you see him grinning.")
							time.Sleep(5 * time.Second)
							fmt.Println("He jumps at you and knocks you over, breaking all your armor.")
							time.Sleep(5 * time.Second)
							fmt.Println("As a last resort you pull out the potion given to you by the shady merchant. You remember his words.")
							time.Sleep(5 * time.Second)
							fmt.Println("Use this to beat the one with the shroom.")
							time.Sleep(5 * time.Second)
							fmt.Println("Toad man marches closer and closer to you and you throw the potion at his head and he disintegrates into ashes.")
							time.Sleep(5 * time.Second)
							fmt.Println("The curse has been lifted, and you hike up the rocky trail. You place your blade next to you as you sit on a rock.")
							time.Sleep(5 * time.Second)
							fmt.Println("Your wolf lays next to you as the two suns set over Mystique Woods.")
							time.Sleep(7 * time.Second)
							for space := 1 ; space <=100 ; space++{
								fmt.Println()
							}
							fmt.Println("THE END")
							for space := 1 ; space <=20 ; space++{
								fmt.Println()
							}
							time.Sleep(4 * time.Second)	
							for space := 1 ; space <=100 ; space++{
								fmt.Println()
							}
							for space := 1 ; space <=100 ; space++{
								fmt.Println()
							}
							fmt.Println("Gerald: Coding")
							for space := 1 ; space <=20 ; space++{
								fmt.Println()
							}
							time.Sleep(4 * time.Second)	
							for space := 1 ; space <=100 ; space++{
								fmt.Println()
							}
							for space := 1 ; space <=100 ; space++{
								fmt.Println()
							}
							fmt.Println("Aarez: Visual and story development")
							for space := 1 ; space <=20 ; space++{
								fmt.Println()
							}
							time.Sleep(4 * time.Second)	
							for space := 1 ; space <=100 ; space++{
								fmt.Println()
							}
							for space := 1 ; space <=100 ; space++{
								fmt.Println()
							}
							fmt.Println("Jonathan: Enemy Info, and Testing")
							for space := 1 ; space <=20 ; space++{
								fmt.Println()
							}
							time.Sleep(4 * time.Second)	
							for space := 1 ; space <=100 ; space++{
								fmt.Println()
							}
							for space := 1 ; space <=100 ; space++{
								fmt.Println()
							}
							fmt.Println("Specail thanks to: Hagan, Suleman, Alex, Zaid, Huzaifa, and Bryce")
							for space := 1 ; space <=20 ; space++{
								fmt.Println()
							}
							time.Sleep(4 * time.Second)	
							for space := 1 ; space <=100 ; space++{
								fmt.Println()
							}
							for space := 1 ; space <=100 ; space++{
								fmt.Println()
							}
							fmt.Println("DEMON'S GATE")
							for space := 1 ; space <=20 ; space++{
								fmt.Println()
							}
							time.Sleep(4 * time.Second)	
							for space := 1 ; space <=100 ; space++{
								fmt.Println()
							}
			} else if game == 2 {
				var a int
				var check = 1
				var b int
				var key = 1
				var key2 = 1
				var key3 = 1
				var mon = 1000000000
				var hp1 = 1000000
				var hp2 = 75
				var hp2or = hp2
				var hp1or = hp1
				var attack1 = 10000000
				var attack2 = 30000000
				var attack3 = 20000000
				var attack4 = 50000000
				var hp1tr int
				var attack1tr int
				var attack2tr int
				var attack3tr int
				var attack4tr int
				var attacken1 = 10
				var attacken2 = 20 
				var attacken3 = 10
				var montr int
				var check4 int
				for {
					attack1 = attack1 - attack1tr
					mon = mon - montr
					hp1 = hp1 - hp1tr
					attack2 = attack2 - attack2tr
					attack3 = attack3 - attack3tr
					attack4 = attack4 - attack4tr
					hp1 = hp1or
					hp2 = 75
					attacken1 = 10
					attacken2 = 20
					attacken3 = 10
					hp2or = hp2
					attack1tr = 0
					hp1tr = 0
					attack2tr = 0
					attack3tr = 0
					attack4tr = 0
					montr = 0
					check = 1
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
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 3 {
									hp2 = hp2 - attack4 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}else {
									hp2 = hp2 - attack4
									fmt.Println("Enemy took",attack4,"damage")
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}
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
						fmt.Println()
						fmt.Println("choose your attack")
						fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
						fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
						fmt.Println("3.heal",attack3,"health 100% accuracy")
						fmt.Scanln(&a)
						if a == 1 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(8)
							if random == 0 {
								hp2 = hp2 - attack1 * 3 / 2
								fmt.Println("Critical Hit")
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}	
							}else {
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy took",attack1,"damage")
								hp2 = hp2-attack1
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}
							}
						} else if a == 2 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(3)
							if random == 1 || random == 2 || random == 3 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 3 {
									hp2 = hp2 - attack2 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}else {
								time.Sleep(1 * time.Second)
								hp2 = hp2 - attack2
								fmt.Println("Enemy took",attack2,"damage")
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}
								}
							} else {
								time.Sleep(1 * time.Second)
								fmt.Println("attack missed")
								time.Sleep(1 * time.Second)
							}
						} else if a == 3 {
							if hp1 == hp1or {
								time.Sleep(1 * time.Second)
								fmt.Println("You're not low enough to heal")
								time.Sleep(1 * time.Second)
							} else if hp1 >= (hp1or - attack3){
								time.Sleep(1 * time.Second)
								hp1 = hp1or
								fmt.Println("You have been healed to max")
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
						time.Sleep(1 * time.Second)
						fmt.Println("enemy is attacking")
						time.Sleep(1 * time.Second)
						if hp2 <= (hp2or / 2) && hp1 > attacken2 {
							time.Sleep(1 * time.Second)
							hp2 = hp2 + attacken3
							fmt.Println("enemy has been healed",attacken3,"health")
							time.Sleep(1 * time.Second)
							continue
						}
						if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(8)
							if random == 0 {
								hp1 = hp1 - attacken1 * 3 / 2
								fmt.Println("Critical Hit")
								time.Sleep(1 * time.Second)
								fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
								time.Sleep(1 * time.Second)
								if hp1 <= 0 {
									fmt.Println("You Died")
									check = 2
									break
								}	
							}else {
								fmt.Println("you took",attacken1,"damage")
								hp1 = hp1-attacken1
								time.Sleep(1 * time.Second)
								if hp1 <= 0 {
									fmt.Println("You Died")
									check = 2
									break
								}
							}
						} else {
							var b = time.Now().UnixNano()
							rand.Seed(b)
							var random3 = rand.Intn(5)
								if random3 == 1 || random3 == 2 || random3 == 3 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp1 = hp1 - attacken2 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											check = 2
											break
										}	
									}else {
										fmt.Println("you took",attacken2,"damage")
										hp1 = hp1-attacken2
										time.Sleep(1 * time.Second)	
										if hp1 <= 0 {
											fmt.Println("You Died")
											check = 2
											break
										}
									}
								} else {
									fmt.Println("attack missed")
										time.Sleep(1 * time.Second)	
										continue
								}
							}
						}
						if check == 2 {
							time.Sleep(2 * time.Second)
							fmt.Println("Respawning")
							time.Sleep(5 * time.Second)
							continue
						}
						time.Sleep(2 * time.Second)
						fmt.Println("The Wood Dweller drops dead, and you are victorious. You scavenge its body and claim its weapon")
						time.Sleep(2 * time.Second)
						attack1 = attack1 + 10
						attack2 = attack2 + 10
						attack4 = attack4 + 10
						attack1tr = attack1tr + 10
						attack2tr = attack2tr + 10
						attack4tr = attack4tr + 10
						fmt.Println("Weapon upgraded")
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
							hp1 = hp1or
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
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 3 {
											hp2 = hp2 - attack4 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
											hp2 = hp2 - attack4
											fmt.Println("Enemy took",attack4,"damage")
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
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
								fmt.Println()
								fmt.Println("choose your attack")
								fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
								fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
								fmt.Println("3.heal",attack3,"health 100% accuracy")
								fmt.Scanln(&a)
								if a == 1 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp2 = hp2 - attack1 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									} else {
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack1,"damage")
										hp2 = hp2-attack1
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
									}
								} else if a == 2 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(3)
									if random == 1 || random == 2 || random == 3 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 3 {
											hp2 = hp2 - attack2 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										} else {
										time.Sleep(1 * time.Second)
										hp2 = hp2 - attack2
										fmt.Println("Enemy took",attack2,"damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
										}
									} else {
										time.Sleep(1 * time.Second)
										fmt.Println("attack missed")
										time.Sleep(1 * time.Second)
									}
								} else if a == 3 {
									if hp1 == hp1or {
										time.Sleep(1 * time.Second)
										fmt.Println("You're not low enough to heal")
										time.Sleep(1 * time.Second)
									} else if hp1 >= (hp1or - attack3){
										time.Sleep(1 * time.Second)
										hp1 = hp1or
										fmt.Println("You have been healed to max")
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
								time.Sleep(1 * time.Second)
								fmt.Println("enemy is attacking")
								time.Sleep(1 * time.Second)
								if hp2 <= (hp2or / 2) && hp1 > attacken2 {
									time.Sleep(1 * time.Second)
									hp2 = hp2 + attacken3
									fmt.Println("enemy has been healed",attacken3,"health")
									time.Sleep(1 * time.Second)
									continue
								}
								if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp1 = hp1 - attacken1 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											check = 2
											break
										}	
									}else {
										fmt.Println("you took",attacken1,"damage")
										hp1 = hp1-attacken1
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											check = 2
											break
										}
									}
								} else {
									var b = time.Now().UnixNano()
									rand.Seed(b)
									var random3 = rand.Intn(5)
										if random3 == 1 || random3 == 2 || random3 == 3 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp1 = hp1 - attacken2 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2
													break
												}	
											}else {
												fmt.Println("you took",attacken2,"damage")
												hp1 = hp1-attacken2
												time.Sleep(1 * time.Second)	
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2
													break
												}
											}
										} else {
											fmt.Println("attack missed")
												time.Sleep(1 * time.Second)	
												continue
										}
									}
								}
							if check == 2 {
								time.Sleep(2 * time.Second)
								fmt.Println("Respawning")
								time.Sleep(5 * time.Second)
								continue
							}
							attack1tr = 0
							hp1tr = 0
							attack2tr = 0
							attack3tr = 0
							attack4tr = 0
							montr = 0
							break
				}
				time.Sleep(1 * time.Second)
					fmt.Println("You gained 20 coins")
					time.Sleep(1 * time.Second)
					mon = mon + 20
					fmt.Println("Total coins:",mon)
					time.Sleep(3 * time.Second)
				for {
					attack1 = attack1 - attack1tr
					mon = mon - montr
					hp1 = hp1 - hp1tr
					attack2 = attack2 - attack2tr
					attack3 = attack3 - attack3tr
					attack4 = attack4 - attack4tr
					hp1 = hp1or
					hp2 = hp2or
					attack1tr = 0
					hp1tr = 0
					attack2tr = 0
					attack3tr = 0
					attack4tr = 0
					montr = 0
					check = 1
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
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 3 {
										hp2 = hp2 - attack4 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									}else {
										hp2 = hp2 - attack4
										fmt.Println("Enemy took",attack4,"damage")
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
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
							fmt.Println()
							fmt.Println("choose your attack")
							fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
							fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
							fmt.Println("3.heal",attack3,"health 100% accuracy")
							fmt.Scanln(&a)
							if a == 1 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 0 {
									hp2 = hp2 - attack1 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}else {
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack1,"damage")
									hp2 = hp2-attack1
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}
								}
							} else if a == 2 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(3)
								if random == 1 || random == 2 || random == 3 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 3 {
										hp2 = hp2 - attack2 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									}else {
									time.Sleep(1 * time.Second)
									hp2 = hp2 - attack2
									fmt.Println("Enemy took",attack2,"damage")
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}
									}
								} else {
									time.Sleep(1 * time.Second)
									fmt.Println("attack missed")
									time.Sleep(1 * time.Second)
								}
							} else if a == 3 {
								if hp1 == hp1or {
									time.Sleep(1 * time.Second)
									fmt.Println("You're not low enough to heal")
									time.Sleep(1 * time.Second)
								} else if hp1 >= (hp1or - attack3){
									time.Sleep(1 * time.Second)
									hp1 = hp1or
									fmt.Println("You have been healed to max")
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
							time.Sleep(1 * time.Second)
							fmt.Println("enemy is attacking")
							time.Sleep(1 * time.Second)
							if hp2 <= (hp2or / 2) && hp1 > attacken2 {
								time.Sleep(1 * time.Second)
								hp2 = hp2 + attacken3
								fmt.Println("enemy has been healed",attacken3,"health")
								time.Sleep(1 * time.Second)
								continue
							}
							if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 0 {
									hp1 = hp1 - attacken1 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
									time.Sleep(1 * time.Second)
									if hp1 <= 0 {
										fmt.Println("You Died")
										check = 2
										break
									}	
								}else {
									fmt.Println("you took",attacken1,"damage")
									hp1 = hp1-attacken1
									time.Sleep(1 * time.Second)
									if hp1 <= 0 {
										fmt.Println("You Died")
										check = 2
										break
									}
								}
							} else {
								var b = time.Now().UnixNano()
								rand.Seed(b)
								var random3 = rand.Intn(5)
									if random3 == 1 || random3 == 2 || random3 == 3 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp1 = hp1 - attacken2 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}	
										}else {
											fmt.Println("you took",attacken2,"damage")
											hp1 = hp1-attacken2
											time.Sleep(1 * time.Second)	
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}
										}
									} else {
										fmt.Println("attack missed")
											time.Sleep(1 * time.Second)	
											continue
									}
								}
							}
						if check == 2 {
							time.Sleep(2 * time.Second)
							fmt.Println("Respawning")
							time.Sleep(5 * time.Second)
							continue
						}
						hp2 = hp2or
						mon = mon + 50
						montr = montr + 50
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
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 3 {
									hp2 = hp2 - attack4 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}else {
									hp2 = hp2 - attack4
									fmt.Println("Enemy took",attack4,"damage")
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}
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
						fmt.Println()
						fmt.Println("choose your attack")
						fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
						fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
						fmt.Println("3.heal",attack3,"health 100% accuracy")
						fmt.Scanln(&a)
						if a == 1 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(8)
							if random == 0 {
								hp2 = hp2 - attack1 * 3 / 2
								fmt.Println("Critical Hit")
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}	
							}else {
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy took",attack1,"damage")
								hp2 = hp2-attack1
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}
							}
						} else if a == 2 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(3)
							if random == 1 || random == 2 || random == 3 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 3 {
									hp2 = hp2 - attack2 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}else {
								time.Sleep(1 * time.Second)
								hp2 = hp2 - attack2
								fmt.Println("Enemy took",attack2,"damage")
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}
								}
							} else {
								time.Sleep(1 * time.Second)
								fmt.Println("attack missed")
								time.Sleep(1 * time.Second)
							}
						} else if a == 3 {
							if hp1 == hp1or {
								time.Sleep(1 * time.Second)
								fmt.Println("You're not low enough to heal")
								time.Sleep(1 * time.Second)
							} else if hp1 >= (hp1or - attack3){
								time.Sleep(1 * time.Second)
								hp1 = hp1or
								fmt.Println("You have been healed to max")
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
						time.Sleep(1 * time.Second)
						fmt.Println("enemy is attacking")
						time.Sleep(1 * time.Second)
						if hp2 <= (hp2or / 2) && hp1 > attacken2 {
							time.Sleep(1 * time.Second)
							hp2 = hp2 + attacken3
							fmt.Println("enemy has been healed",attacken3,"health")
							time.Sleep(1 * time.Second)
							continue
						}
						if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(8)
							if random == 0 {
								hp1 = hp1 - attacken1 * 3 / 2
								fmt.Println("Critical Hit")
								time.Sleep(1 * time.Second)
								fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
								time.Sleep(1 * time.Second)
								if hp1 <= 0 {
									fmt.Println("You Died")
									check = 2
									break
								}	
							}else {
								fmt.Println("you took",attacken1,"damage")
								hp1 = hp1-attacken1
								time.Sleep(1 * time.Second)
								if hp1 <= 0 {
									fmt.Println("You Died")
									check = 2
									break
								}
							}
						} else {
							var b = time.Now().UnixNano()
							rand.Seed(b)
							var random3 = rand.Intn(5)
								if random3 == 1 || random3 == 2 || random3 == 3 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp1 = hp1 - attacken2 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											check = 2
											break
										}	
									}else {
										fmt.Println("you took",attacken2,"damage")
										hp1 = hp1-attacken2
										time.Sleep(1 * time.Second)	
										if hp1 <= 0 {
											fmt.Println("You Died")
											check = 2
											break
										}
									}
								} else {
									fmt.Println("attack missed")
										time.Sleep(1 * time.Second)	
										continue
								}
							}
						}
					if check == 2 {
						time.Sleep(2 * time.Second)
						fmt.Println("Respawning")
						time.Sleep(5 * time.Second)
						continue
					}
					attack1tr = 0
					hp1tr = 0
					attack2tr = 0
					attack3tr = 0
					attack4tr = 0
					montr = 0
					hp1 = hp1or
					mon = mon + 30
					time.Sleep(1 * time.Second)
						fmt.Println("You gained 30 coins")
						time.Sleep(1 * time.Second)
					fmt.Println("Total coins:",mon)
					time.Sleep(4 * time.Second)
					break
				}
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
						hp1 = hp1or
						hp2 = hp2or
						check = 1
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
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 3 {
													hp2 = hp2 - attack4 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}else {
													hp2 = hp2 - attack4
													fmt.Println("Enemy took",attack4,"damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}
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
										fmt.Println()
										fmt.Println("choose your attack")
										fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
										fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
										fmt.Println("3.heal",attack3,"health 100% accuracy")
										fmt.Scanln(&a)
										if a == 1 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp2 = hp2 - attack1 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack1,"damage")
												hp2 = hp2-attack1
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
											}
										} else if a == 2 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(3)
											if random == 1 || random == 2 || random == 3 {
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 3 {
													hp2 = hp2 - attack2 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
													time.Sleep(1 * time.Second)
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}else {
												time.Sleep(1 * time.Second)
												hp2 = hp2 - attack2
												fmt.Println("Enemy took",attack2,"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
												}
											} else {
												time.Sleep(1 * time.Second)
												fmt.Println("attack missed")
												time.Sleep(1 * time.Second)
											}
										} else if a == 3 {
											if hp1 == hp1or {
												time.Sleep(1 * time.Second)
												fmt.Println("You're not low enough to heal")
												time.Sleep(1 * time.Second)
											} else if hp1 >= (hp1or - attack3){
												time.Sleep(1 * time.Second)
												hp1 = hp1or
												fmt.Println("You have been healed to max")
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
										time.Sleep(1 * time.Second)
										fmt.Println("enemy is attacking")
										time.Sleep(1 * time.Second)
										if hp2 <= (hp2or / 2) && hp1 > attacken2 {
											time.Sleep(1 * time.Second)
											hp2 = hp2 + attacken3
											fmt.Println("enemy has been healed",attacken3,"health")
											time.Sleep(1 * time.Second)
											continue
										}
										if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp1 = hp1 - attacken1 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2
													break
												}	
											}else {
												fmt.Println("you took",attacken1,"damage")
												hp1 = hp1-attacken1
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2
													break
												}
											}
										} else {
											var b = time.Now().UnixNano()
											rand.Seed(b)
											var random3 = rand.Intn(5)
												if random3 == 1 || random3 == 2 || random3 == 3 {
													var s = time.Now().UnixNano()
													rand.Seed(s)
													var random = rand.Intn(8)
													if random == 0 {
														hp1 = hp1 - attacken2 * 3 / 2
														fmt.Println("Critical Hit")
														time.Sleep(1 * time.Second)
														fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
														time.Sleep(1 * time.Second)
														if hp1 <= 0 {
															fmt.Println("You Died")
															check = 2
															break
														}	
													}else {
														fmt.Println("you took",attacken2,"damage")
														hp1 = hp1-attacken2
														time.Sleep(1 * time.Second)	
														if hp1 <= 0 {
															fmt.Println("You Died")
															check = 2
															break
														}
													}
												} else {
													fmt.Println("attack missed")
														time.Sleep(1 * time.Second)	
														continue
												}
											}
										}
										if check == 2 {
											time.Sleep(2 * time.Second)
											fmt.Println("Respawning")
											time.Sleep(5 * time.Second)
											key = 1
											continue
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
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 3 {
													hp2 = hp2 - attack4 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}else {
													hp2 = hp2 - attack4
													fmt.Println("Enemy took",attack4,"damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}
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
										fmt.Println()
										fmt.Println("choose your attack")
										fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
										fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
										fmt.Println("3.heal",attack3,"health 100% accuracy")
										fmt.Scanln(&a)
										if a == 1 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp2 = hp2 - attack1 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack1,"damage")
												hp2 = hp2-attack1
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
											}
										} else if a == 2 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(3)
											if random == 1 || random == 2 || random == 3 {
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 3 {
													hp2 = hp2 - attack2 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
													time.Sleep(1 * time.Second)
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}else {
												time.Sleep(1 * time.Second)
												hp2 = hp2 - attack2
												fmt.Println("Enemy took",attack2,"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
												}
											} else {
												time.Sleep(1 * time.Second)
												fmt.Println("attack missed")
												time.Sleep(1 * time.Second)
											}
										} else if a == 3 {
											if hp1 == hp1or {
												time.Sleep(1 * time.Second)
												fmt.Println("You're not low enough to heal")
												time.Sleep(1 * time.Second)
											} else if hp1 >= (hp1or - attack3){
												time.Sleep(1 * time.Second)
												hp1 = hp1or
												fmt.Println("You have been healed to max")
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
										time.Sleep(1 * time.Second)
										fmt.Println("enemy is attacking")
										time.Sleep(1 * time.Second)
										if hp2 <= (hp2or / 2) && hp1 > attacken2 {
											time.Sleep(1 * time.Second)
											hp2 = hp2 + attacken3
											fmt.Println("enemy has been healed",attacken3,"health")
											time.Sleep(1 * time.Second)
											continue
										}
										if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp1 = hp1 - attacken1 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2
													break
												}	
											}else {
												fmt.Println("you took",attacken1,"damage")
												hp1 = hp1-attacken1
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2
													break
												}
											}
										} else {
											var b = time.Now().UnixNano()
											rand.Seed(b)
											var random3 = rand.Intn(5)
												if random3 == 1 || random3 == 2 || random3 == 3 {
													var s = time.Now().UnixNano()
													rand.Seed(s)
													var random = rand.Intn(8)
													if random == 0 {
														hp1 = hp1 - attacken2 * 3 / 2
														fmt.Println("Critical Hit")
														time.Sleep(1 * time.Second)
														fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
														time.Sleep(1 * time.Second)
														if hp1 <= 0 {
															fmt.Println("You Died")
															check = 2
															break
														}	
													}else {
														fmt.Println("you took",attacken2,"damage")
														hp1 = hp1-attacken2
														time.Sleep(1 * time.Second)	
														if hp1 <= 0 {
															fmt.Println("You Died")
															check = 2
															break
														}
													}
												} else {
													fmt.Println("attack missed")
														time.Sleep(1 * time.Second)	
														continue
												}
											}
										}
									if check == 2 {
										time.Sleep(2 * time.Second)
										fmt.Println("Respawning")
										time.Sleep(5 * time.Second)
										key2 = 1
										continue
									}
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
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 3 {
													hp2 = hp2 - attack4 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}else {
													hp2 = hp2 - attack4
													fmt.Println("Enemy took",attack4,"damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}
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
										fmt.Println()
										fmt.Println("choose your attack")
										fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
										fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
										fmt.Println("3.heal",attack3,"health 100% accuracy")
										fmt.Scanln(&a)
										if a == 1 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp2 = hp2 - attack1 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack1,"damage")
												hp2 = hp2-attack1
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
											}
										} else if a == 2 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(3)
											if random == 1 || random == 2 || random == 3 {
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 3 {
													hp2 = hp2 - attack2 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
													time.Sleep(1 * time.Second)
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}else {
												time.Sleep(1 * time.Second)
												hp2 = hp2 - attack2
												fmt.Println("Enemy took",attack2,"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
												}
											} else {
												time.Sleep(1 * time.Second)
												fmt.Println("attack missed")
												time.Sleep(1 * time.Second)
											}
										} else if a == 3 {
											if hp1 == hp1or {
												time.Sleep(1 * time.Second)
												fmt.Println("You're not low enough to heal")
												time.Sleep(1 * time.Second)
											} else if hp1 >= (hp1or - attack3){
												time.Sleep(1 * time.Second)
												hp1 = hp1or
												fmt.Println("You have been healed to max")
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
										time.Sleep(1 * time.Second)
										fmt.Println("enemy is attacking")
										time.Sleep(1 * time.Second)
										if hp2 <= (hp2or / 2) && hp1 > attacken2 {
											time.Sleep(1 * time.Second)
											hp2 = hp2 + attacken3
											fmt.Println("enemy has been healed",attacken3,"health")
											time.Sleep(1 * time.Second)
											continue
										}
										if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp1 = hp1 - attacken1 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2
													break
												}	
											}else {
												fmt.Println("you took",attacken1,"damage")
												hp1 = hp1-attacken1
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2
													break
												}
											}
										} else {
											var b = time.Now().UnixNano()
											rand.Seed(b)
											var random3 = rand.Intn(5)
												if random3 == 1 || random3 == 2 || random3 == 3 {
													var s = time.Now().UnixNano()
													rand.Seed(s)
													var random = rand.Intn(8)
													if random == 0 {
														hp1 = hp1 - attacken2 * 3 / 2
														fmt.Println("Critical Hit")
														time.Sleep(1 * time.Second)
														fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
														time.Sleep(1 * time.Second)
														if hp1 <= 0 {
															fmt.Println("You Died")
															check = 2
															break
														}	
													}else {
														fmt.Println("you took",attacken2,"damage")
														hp1 = hp1-attacken2
														time.Sleep(1 * time.Second)	
														if hp1 <= 0 {
															fmt.Println("You Died")
															check = 2
															break
														}
													}
												} else {
													fmt.Println("attack missed")
														time.Sleep(1 * time.Second)	
														continue
												}
											}
										}
									if check == 2 {
										time.Sleep(2 * time.Second)
										fmt.Println("Respawning")
										time.Sleep(5 * time.Second)
										key3 = 1
										continue
									}
									hp1 = hp1or
									mon = mon + 50
									time.Sleep(1 * time.Second)
								fmt.Println("You gained 50 coins")
								}else{
									fmt.Println("You already explored here. departing in 3 seconds")
									time.Sleep(3 * time.Second)
								}
							}else if a == 4 {
								time.Sleep(1 * time.Second)
									fmt.Println("You enter a dungeon with a large monster weilding a giant axe")
									time.Sleep(1 * time.Second)
									fmt.Println("You enter a battle with Taurus")
									hp2 = hp2or
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
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 3 {
													hp2 = hp2 - attack4 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}else {
													hp2 = hp2 - attack4
													fmt.Println("Enemy took",attack4,"damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}
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
										fmt.Println()
										fmt.Println("choose your attack")
										fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
										fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
										fmt.Println("3.heal",attack3,"health 100% accuracy")
										fmt.Scanln(&a)
										if a == 1 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp2 = hp2 - attack1 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack1,"damage")
												hp2 = hp2-attack1
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
											}
										} else if a == 2 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(3)
											if random == 1 || random == 2 || random == 3 {
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 3 {
													hp2 = hp2 - attack2 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
													time.Sleep(1 * time.Second)
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}else {
												time.Sleep(1 * time.Second)
												hp2 = hp2 - attack2
												fmt.Println("Enemy took",attack2,"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
												}
											} else {
												time.Sleep(1 * time.Second)
												fmt.Println("attack missed")
												time.Sleep(1 * time.Second)
											}
										} else if a == 3 {
											if hp1 == hp1or {
												time.Sleep(1 * time.Second)
												fmt.Println("You're not low enough to heal")
												time.Sleep(1 * time.Second)
											} else if hp1 >= (hp1or - attack3){
												time.Sleep(1 * time.Second)
												hp1 = hp1or
												fmt.Println("You have been healed to max")
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
										time.Sleep(1 * time.Second)
										fmt.Println("enemy is attacking")
										time.Sleep(1 * time.Second)
										if hp2 <= (hp2or / 2) && hp1 > attacken2 {
											time.Sleep(1 * time.Second)
											hp2 = hp2 + attacken3
											fmt.Println("enemy has been healed",attacken3,"health")
											time.Sleep(1 * time.Second)
											continue
										}
										if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp1 = hp1 - attacken1 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2
													break
												}	
											}else {
												fmt.Println("you took",attacken1,"damage")
												hp1 = hp1-attacken1
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2
													break
												}
											}
										} else {
											var b = time.Now().UnixNano()
											rand.Seed(b)
											var random3 = rand.Intn(5)
												if random3 == 1 || random3 == 2 || random3 == 3 {
													var s = time.Now().UnixNano()
													rand.Seed(s)
													var random = rand.Intn(8)
													if random == 0 {
														hp1 = hp1 - attacken2 * 3 / 2
														fmt.Println("Critical Hit")
														time.Sleep(1 * time.Second)
														fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
														time.Sleep(1 * time.Second)
														if hp1 <= 0 {
															fmt.Println("You Died")
															check = 2
															break
														}	
													}else {
														fmt.Println("you took",attacken2,"damage")
														hp1 = hp1-attacken2
														time.Sleep(1 * time.Second)	
														if hp1 <= 0 {
															fmt.Println("You Died")
															check = 2
															break
														}
													}
												} else {
													fmt.Println("attack missed")
														time.Sleep(1 * time.Second)	
														continue
												}
											}
										}
									if check == 2 {
										time.Sleep(2 * time.Second)
										fmt.Println("Respawning")
										time.Sleep(5 * time.Second)
										continue
									}
									hp1 = hp1or
									time.Sleep(4 * time.Second)
									mon = mon + 200
									time.Sleep(1 * time.Second)
								fmt.Println("You gained 200 coins")
								time.Sleep(1 * time.Second)
								fmt.Println("Total Coins:",mon)
								time.Sleep(1 * time.Second)
									fmt.Println("After beating the Taurus, you scavenge its carcass, and find a key, which opens the chest.")
									time.Sleep(4 * time.Second)
									fmt.Println("The first, bright yellow gem is inside. You have saved Mystic Woods from the curse.")
									time.Sleep(4 * time.Second)
									fmt.Println("Now, you must go, and conquer Mount Olympus")
								time.Sleep(7 * time.Second)
								check4 = 1
								break
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
									if check == 2 {
								time.Sleep(2 * time.Second)
								fmt.Println("Respawning")
								time.Sleep(5 * time.Second)
								continue
							}
							if a >= 4 || a == 0{
								continue
							}
							if check4 == 1 {
								break
							}
							attack1tr = 0
							hp1tr = 0
							attack2tr = 0
							attack3tr = 0
							attack4tr = 0
							montr = 0
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
					for {
					hp1 = hp1 - hp1tr 
					hp1 = hp1or
					attack1 = attack1 - attack1tr
					mon = mon - montr
					hp1 = hp1 - hp1tr
					hp1or = hp1
					attack2 = attack2 - attack2tr
					attack3 = attack3 - attack3tr
					attack4 = attack4 - attack4tr
					hp2or = hp2
					attack1tr = 0
					hp1tr = 0
					attack2tr = 0
					attack3tr = 0
					attack4tr = 0
					montr = 0
					check = 1
						attacken1 = 20
						attacken2 = 40
						hp2 = 130
						hp2or = hp2
						attacken3 = 30
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
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 3 {
										hp2 = hp2 - attack4 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									}else {
										hp2 = hp2 - attack4
										fmt.Println("Enemy took",attack4,"damage")
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
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
							fmt.Println()
							fmt.Println("choose your attack")
							fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
							fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
							fmt.Println("3.heal",attack3,"health 100% accuracy")
							fmt.Scanln(&a)
							if a == 1 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 0 {
									hp2 = hp2 - attack1 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}else {
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack1,"damage")
									hp2 = hp2-attack1
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}
								}
							} else if a == 2 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(3)
								if random == 1 || random == 2 || random == 3 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 3 {
										hp2 = hp2 - attack2 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									}else {
									time.Sleep(1 * time.Second)
									hp2 = hp2 - attack2
									fmt.Println("Enemy took",attack2,"damage")
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}
									}
								} else {
									time.Sleep(1 * time.Second)
									fmt.Println("attack missed")
									time.Sleep(1 * time.Second)
								}
							} else if a == 3 {
								if hp1 == hp1or {
									time.Sleep(1 * time.Second)
									fmt.Println("You're not low enough to heal")
									time.Sleep(1 * time.Second)
								} else if hp1 >= (hp1or - attack3){
									time.Sleep(1 * time.Second)
									hp1 = hp1or
									fmt.Println("You have been healed to max")
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
							time.Sleep(1 * time.Second)
							fmt.Println("enemy is attacking")
							time.Sleep(1 * time.Second)
							if hp2 <= (hp2or / 2) && hp1 > attacken2 {
								time.Sleep(1 * time.Second)
								hp2 = hp2 + attacken3
								fmt.Println("enemy has been healed",attacken3,"health")
								time.Sleep(1 * time.Second)
								continue
							}
							if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 0 {
									hp1 = hp1 - attacken1 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
									time.Sleep(1 * time.Second)
									if hp1 <= 0 {
										fmt.Println("You Died")
										check = 2
										break
									}	
								}else {
									fmt.Println("you took",attacken1,"damage")
									hp1 = hp1-attacken1
									time.Sleep(1 * time.Second)
									if hp1 <= 0 {
										fmt.Println("You Died")
										check = 2
										break
									}
								}
							} else {
								var b = time.Now().UnixNano()
								rand.Seed(b)
								var random3 = rand.Intn(5)
									if random3 == 1 || random3 == 2 || random3 == 3 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp1 = hp1 - attacken2 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}	
										}else {
											fmt.Println("you took",attacken2,"damage")
											hp1 = hp1-attacken2
											time.Sleep(1 * time.Second)	
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}
										}
									} else {
										fmt.Println("attack missed")
											time.Sleep(1 * time.Second)	
											continue
									}
								}
							}
						if check == 2 {
							time.Sleep(2 * time.Second)
							fmt.Println("Respawning")
							time.Sleep(5 * time.Second)
							continue
						}
						hp1 = hp1or
						hp1 = hp1 + 30
						hp1tr = hp1tr + 30
						hp1or = hp1
						time.Sleep(1 * time.Second)
						fmt.Println("After the battle, you claim its helmet.")
						time.Sleep(4 * time.Second)
						fmt.Println("Armour upgraded")
						time.Sleep(4 * time.Second)
						fmt.Println("You know what's coming for you. There is no good left in this Mountain range. You are aware that you must conquer two other mountains before Olympus.")
						time.Sleep(4 * time.Second)
						fmt.Println("The first mountain, is merely  just a hill. On the top of the mountain, waits the Brood Mother, a lethally dangerous spider.")
						time.Sleep(4 * time.Second)
						fmt.Println("When you hike up the hill, you notice some movement. It is a possessed spartan.  As soon as you notice it, the spartan gets impaled by a ten-foot spider.")
						time.Sleep(4 * time.Second)
						fmt.Println("You run away into a bush, hiding. The spider notices you.")
						hp2 = 190
						hp2or = hp2
						attacken1 = 30
						attacken2 = 50
						attacken3 = 30
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
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 3 {
										hp2 = hp2 - attack4 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									}else {
										hp2 = hp2 - attack4
										fmt.Println("Enemy took",attack4,"damage")
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
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
							fmt.Println()
							fmt.Println("choose your attack")
							fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
							fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
							fmt.Println("3.heal",attack3,"health 100% accuracy")
							fmt.Scanln(&a)
							if a == 1 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 0 {
									hp2 = hp2 - attack1 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}else {
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack1,"damage")
									hp2 = hp2-attack1
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}
								}
							} else if a == 2 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(3)
								if random == 1 || random == 2 || random == 3 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 3 {
										hp2 = hp2 - attack2 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									}else {
									time.Sleep(1 * time.Second)
									hp2 = hp2 - attack2
									fmt.Println("Enemy took",attack2,"damage")
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}
									}
								} else {
									time.Sleep(1 * time.Second)
									fmt.Println("attack missed")
									time.Sleep(1 * time.Second)
								}
							} else if a == 3 {
								if hp1 == hp1or {
									time.Sleep(1 * time.Second)
									fmt.Println("You're not low enough to heal")
									time.Sleep(1 * time.Second)
								} else if hp1 >= (hp1or - attack3){
									time.Sleep(1 * time.Second)
									hp1 = hp1or
									fmt.Println("You have been healed to max")
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
							time.Sleep(1 * time.Second)
							fmt.Println("enemy is attacking")
							time.Sleep(1 * time.Second)
							if hp2 <= (hp2or / 2) && hp1 > attacken2 {
								time.Sleep(1 * time.Second)
								hp2 = hp2 + attacken3
								fmt.Println("enemy has been healed",attacken3,"health")
								time.Sleep(1 * time.Second)
								continue
							}
							if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 0 {
									hp1 = hp1 - attacken1 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
									time.Sleep(1 * time.Second)
									if hp1 <= 0 {
										fmt.Println("You Died")
										check = 2
										break
									}	
								}else {
									fmt.Println("you took",attacken1,"damage")
									hp1 = hp1-attacken1
									time.Sleep(1 * time.Second)
									if hp1 <= 0 {
										fmt.Println("You Died")
										check = 2
										break
									}
								}
							} else {
								var b = time.Now().UnixNano()
								rand.Seed(b)
								var random3 = rand.Intn(5)
									if random3 == 1 || random3 == 2 || random3 == 3 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp1 = hp1 - attacken2 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}	
										}else {
											fmt.Println("you took",attacken2,"damage")
											hp1 = hp1-attacken2
											time.Sleep(1 * time.Second)	
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}
										}
									} else {
										fmt.Println("attack missed")
											time.Sleep(1 * time.Second)	
											continue
									}
								}
							}
							if check == 2 {
								time.Sleep(2 * time.Second)
								fmt.Println("Respawning")
								time.Sleep(5 * time.Second)
								continue
							}
							attack1tr = 0
							hp1tr = 0
							attack2tr = 0
							attack3tr = 0
							attack4tr = 0
							montr = 0
							break
					}
					for {
					hp1 = hp1 - hp1tr 
					hp1 = hp1or
					attack1 = attack1 - attack1tr
					mon = mon - montr
					hp1 = hp1 - hp1tr
					hp1or = hp1
					attack2 = attack2 - attack2tr
					attack3 = attack3 - attack3tr
					attack4 = attack4 - attack4tr
					hp2or = hp2
					attack1tr = 0
					hp1tr = 0
					attack2tr = 0
					attack3tr = 0
					attack4tr = 0
					montr = 0
					check = 1
					time.Sleep(4 * time.Second)
					fmt.Println("The Brood Mother flees, and you continue up the hill. You hear a creak in the distance, Do you go there:")
					time.Sleep(4 * time.Second)
					fmt.Println("1.Yes")
					fmt.Println("2.No")
					fmt.Scanln(&b)
					if b == 1 {
						fmt.Println("You descend the hill, but you trip, fall and die.")
						time.Sleep(4 * time.Second)
						fmt.Println("Respawning")
						time.Sleep(5 * time.Second)
						continue
					}
					fmt.Println("You have just avoided a very disappointing death. You nervously advance forward.")
					time.Sleep(4 * time.Second)
					fmt.Println("you have reached the top of the hill. but you don't see the spider. What do you do? Ring the bell, or run back")
					time.Sleep(4 * time.Second)
					fmt.Println("1.Ring bell")
					fmt.Println("2.Run back")
					fmt.Scanln(&b)
					if b == 2 {
						fmt.Println("You run down the hill, scared. you feel a sharp pain in the back.")
						time.Sleep(4 * time.Second)
						fmt.Println("Just like the possessed Spartan, you have been impaled. Dead.")
						time.Sleep( 4 * time.Second)
						fmt.Println("Respawning")
						time.Sleep( 5 * time.Second)
						continue
					}
					time.Sleep(4 * time.Second)
					fmt.Println("You ring the bell, you have effortlessly conquered the first Mountain.")
					time.Sleep(4 * time.Second)
					fmt.Println("You advance towards the second Mountain, which is way larger than the first, but way smaller than Mt. Olympus.")
					time.Sleep(4 * time.Second)
					fmt.Println("You are stopped in your path. It is the Brood Mother.")
					hp1 = hp1or
					hp2 = 100
					hp2or = hp2
					attacken1 = 20
					attacken2 = 40
					attacken3 = 30
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
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 3 {
									hp2 = hp2 - attack4 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}else {
									hp2 = hp2 - attack4
									fmt.Println("Enemy took",attack4,"damage")
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}
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
						fmt.Println()
						fmt.Println("choose your attack")
						fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
						fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
						fmt.Println("3.heal",attack3,"health 100% accuracy")
						fmt.Scanln(&a)
						if a == 1 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(8)
							if random == 0 {
								hp2 = hp2 - attack1 * 3 / 2
								fmt.Println("Critical Hit")
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}	
							}else {
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy took",attack1,"damage")
								hp2 = hp2-attack1
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}
							}
						} else if a == 2 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(3)
							if random == 1 || random == 2 || random == 3 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 3 {
									hp2 = hp2 - attack2 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}else {
								time.Sleep(1 * time.Second)
								hp2 = hp2 - attack2
								fmt.Println("Enemy took",attack2,"damage")
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}
								}
							} else {
								time.Sleep(1 * time.Second)
								fmt.Println("attack missed")
								time.Sleep(1 * time.Second)
							}
						} else if a == 3 {
							if hp1 == hp1or {
								time.Sleep(1 * time.Second)
								fmt.Println("You're not low enough to heal")
								time.Sleep(1 * time.Second)
							} else if hp1 >= (hp1or - attack3){
								time.Sleep(1 * time.Second)
								hp1 = hp1or
								fmt.Println("You have been healed to max")
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
						time.Sleep(1 * time.Second)
						fmt.Println("enemy is attacking")
						time.Sleep(1 * time.Second)
						if hp2 <= (hp2or / 2) && hp1 > attacken2 {
							time.Sleep(1 * time.Second)
							hp2 = hp2 + attacken3
							fmt.Println("enemy has been healed",attacken3,"health")
							time.Sleep(1 * time.Second)
							continue
						}
						if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(8)
							if random == 0 {
								hp1 = hp1 - attacken1 * 3 / 2
								fmt.Println("Critical Hit")
								time.Sleep(1 * time.Second)
								fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
								time.Sleep(1 * time.Second)
								if hp1 <= 0 {
									fmt.Println("You Died")
									check = 2
									break
								}	
							}else {
								fmt.Println("you took",attacken1,"damage")
								hp1 = hp1-attacken1
								time.Sleep(1 * time.Second)
								if hp1 <= 0 {
									fmt.Println("You Died")
									check = 2
									break
								}
							}
						} else {
							var b = time.Now().UnixNano()
							rand.Seed(b)
							var random3 = rand.Intn(5)
								if random3 == 1 || random3 == 2 || random3 == 3 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp1 = hp1 - attacken2 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											check = 2
											break
										}	
									}else {
										fmt.Println("you took",attacken2,"damage")
										hp1 = hp1-attacken2
										time.Sleep(1 * time.Second)	
										if hp1 <= 0 {
											fmt.Println("You Died")
											check = 2
											break
										}
									}
								} else {
									fmt.Println("attack missed")
										time.Sleep(1 * time.Second)	
										continue
								}
							}
						}
						if check == 2 {
							time.Sleep(2 * time.Second)
							fmt.Println("Respawning")
							time.Sleep(5 * time.Second)
							continue
						}
					hp1 = hp1or
					attack1tr = 0
					hp1tr = 0
					attack2tr = 0
					attack3tr = 0
					attack4tr = 0
					montr = 0
					break
					}
							attack3 = attack3 + 20
							fmt.Println("The Brood Mother was powerless. You had reduced her to cripples in the first battle.")
							time.Sleep(4 * time.Second)
							fmt.Println("You scavenge her body, and find an advanced heal potion.")
							time.Sleep(4 * time.Second)
							fmt.Println("Magic upgraded")
							time.Sleep(4 * time.Second)
							fmt.Println("The second mountain, home of the Mountain beast, who has low health, but packs a great punch.")
							time.Sleep(4 * time.Second)
							fmt.Println("He guards the second bell, so you advance up the mountain. on the path, you encounter a Lost Triton.")
							time.Sleep(4 * time.Second)
							fmt.Println("A minion which wields a trident.")
							for {
								hp1 = hp1 - hp1tr 
								hp1 = hp1or
								attack1 = attack1 - attack1tr
								mon = mon - montr
								hp1 = hp1 - hp1tr
								hp1or = hp1
								attack2 = attack2 - attack2tr
								attack3 = attack3 - attack3tr
								attack4 = attack4 - attack4tr
								hp2or = hp2
								attack1tr = 0
								hp1tr = 0
								attack2tr = 0
								attack3tr = 0
								attack4tr = 0
								montr = 0
								check = 1
								hp2 = 130
								hp2or = hp2
								attacken1 = 35
								attacken2 = 55
								attacken3 = 30
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
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack4 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
												hp2 = hp2 - attack4
												fmt.Println("Enemy took",attack4,"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
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
									fmt.Println()
									fmt.Println("choose your attack")
									fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
									fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
									fmt.Println("3.heal",attack3,"health 100% accuracy")
									fmt.Scanln(&a)
									if a == 1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp2 = hp2 - attack1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1,"damage")
											hp2 = hp2-attack1
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
										}
									} else if a == 2 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(3)
										if random == 1 || random == 2 || random == 3 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack2 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
											time.Sleep(1 * time.Second)
											hp2 = hp2 - attack2
											fmt.Println("Enemy took",attack2,"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
											}
										} else {
											time.Sleep(1 * time.Second)
											fmt.Println("attack missed")
											time.Sleep(1 * time.Second)
										}
									} else if a == 3 {
										if hp1 == hp1or {
											time.Sleep(1 * time.Second)
											fmt.Println("You're not low enough to heal")
											time.Sleep(1 * time.Second)
										} else if hp1 >= (hp1or - attack3){
											time.Sleep(1 * time.Second)
											hp1 = hp1or
											fmt.Println("You have been healed to max")
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
									time.Sleep(1 * time.Second)
									fmt.Println("enemy is attacking")
									time.Sleep(1 * time.Second)
									if hp2 <= (hp2or / 2) && hp1 > attacken2 {
										time.Sleep(1 * time.Second)
										hp2 = hp2 + attacken3
										fmt.Println("enemy has been healed",attacken3,"health")
										time.Sleep(1 * time.Second)
										continue
									}
									if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp1 = hp1 - attacken1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}	
										}else {
											fmt.Println("you took",attacken1,"damage")
											hp1 = hp1-attacken1
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}
										}
									} else {
										var b = time.Now().UnixNano()
										rand.Seed(b)
										var random3 = rand.Intn(5)
											if random3 == 1 || random3 == 2 || random3 == 3 {
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 0 {
													hp1 = hp1 - attacken2 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
													time.Sleep(1 * time.Second)
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2
														break
													}	
												}else {
													fmt.Println("you took",attacken2,"damage")
													hp1 = hp1-attacken2
													time.Sleep(1 * time.Second)	
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2
														break
													}
												}
											} else {
												fmt.Println("attack missed")
													time.Sleep(1 * time.Second)	
													continue
											}
										}
									}
									if check == 2 {
										time.Sleep(2 * time.Second)
										fmt.Println("Respawning")
										time.Sleep(5 * time.Second)
										continue
									}
								hp1 = hp1or
								time.Sleep(4 * time.Second)
								fmt.Println("After beating the Triton, you notice it dropped something, do you ignore it or pick it up ?")
								time.Sleep(4 * time.Second)
								fmt.Println("1.Pick it up.")
								fmt.Println("2.Ignore it.")
								fmt.Scanln(&b)
								if b == 1 {
									fmt.Println("New Weapon: Trident")
									attack1tr = attack1
									attack2tr = attack2
									attack4tr = attack4
									attack2 = 75000000
									attack1 = 50000000
									attack4 = 100000000
									attack1tr = attack1 - attack1tr
									attack2tr = attack2 - attack2tr
									attack4tr = attack4 - attack4tr
								}
								time.Sleep(4 * time.Second)
								fmt.Println("you proceed further up, until you see a wall, too hard to climb.")
								time.Sleep(4 * time.Second)
								fmt.Println("You attempt to climb the wall, and notice that you are at a plateau, almost up the mountain. You encounter a demon.")
								time.Sleep(4 * time.Second)
								fmt.Println("It is one of Hades' minions.")
								time.Sleep(4 * time.Second)
								hp2 = 180
								hp2or = hp2
								attacken1 = 35
								attacken2 = 65
								attacken3 = 30
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
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack4 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
												hp2 = hp2 - attack4
												fmt.Println("Enemy took",attack4,"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
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
									fmt.Println()
									fmt.Println("choose your attack")
									fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
									fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
									fmt.Println("3.heal",attack3,"health 100% accuracy")
									fmt.Scanln(&a)
									if a == 1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp2 = hp2 - attack1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1,"damage")
											hp2 = hp2-attack1
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
										}
									} else if a == 2 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(3)
										if random == 1 || random == 2 || random == 3 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack2 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
											time.Sleep(1 * time.Second)
											hp2 = hp2 - attack2
											fmt.Println("Enemy took",attack2,"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
											}
										} else {
											time.Sleep(1 * time.Second)
											fmt.Println("attack missed")
											time.Sleep(1 * time.Second)
										}
									} else if a == 3 {
										if hp1 == hp1or {
											time.Sleep(1 * time.Second)
											fmt.Println("You're not low enough to heal")
											time.Sleep(1 * time.Second)
										} else if hp1 >= (hp1or - attack3){
											time.Sleep(1 * time.Second)
											hp1 = hp1or
											fmt.Println("You have been healed to max")
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
									time.Sleep(1 * time.Second)
									fmt.Println("enemy is attacking")
									time.Sleep(1 * time.Second)
									if hp2 <= (hp2or / 2) && hp1 > attacken2 {
										time.Sleep(1 * time.Second)
										hp2 = hp2 + attacken3
										fmt.Println("enemy has been healed",attacken3,"health")
										time.Sleep(1 * time.Second)
										continue
									}
									if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp1 = hp1 - attacken1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}	
										}else {
											fmt.Println("you took",attacken1,"damage")
											hp1 = hp1-attacken1
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}
										}
									} else {
										var b = time.Now().UnixNano()
										rand.Seed(b)
										var random3 = rand.Intn(5)
											if random3 == 1 || random3 == 2 || random3 == 3 {
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 0 {
													hp1 = hp1 - attacken2 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
													time.Sleep(1 * time.Second)
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2
														break
													}	
												}else {
													fmt.Println("you took",attacken2,"damage")
													hp1 = hp1-attacken2
													time.Sleep(1 * time.Second)	
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2
														break
													}
												}
											} else {
												fmt.Println("attack missed")
													time.Sleep(1 * time.Second)	
													continue
											}
										}
									}
									if check == 2 {
										time.Sleep(2 * time.Second)
										fmt.Println("Respawning")
										time.Sleep(5 * time.Second)
										continue
									}
								hp1 = hp1or
								attack1tr = 0
								hp1tr = 0
								attack2tr = 0
								attack3tr = 0
								attack4tr = 0
								montr = 0
								break
							}
						hp1 = hp1 + 20
						hp1or = hp1
						for {
							fmt.Println("After beating the minion, your head feels a little relieved.")
							time.Sleep(4 * time.Second)
							fmt.Println("You pick up his helmet")
							time.Sleep(4 * time.Second)
							fmt.Println("Armour upgraded")
							time.Sleep(4 * time.Second)
							fmt.Println("This means that if you get rid of their master, it undo's the curse. You continue up the hill, and reach the top.")
							time.Sleep(4 * time.Second)
							fmt.Println("You see a Skeletal Giant. You must beat him to ring the second bell. You confidently enter the battle.")
							time.Sleep(4 * time.Second)
							hp1 = hp1 - hp1tr 
							hp1 = hp1or
							attack1 = attack1 - attack1tr
							mon = mon - montr
							hp1 = hp1 - hp1tr
							hp1or = hp1
							attack2 = attack2 - attack2tr
							attack3 = attack3 - attack3tr
							attack4 = attack4 - attack4tr
							hp2or = hp2
							attack1tr = 0
							hp1tr = 0
							attack2tr = 0
							attack3tr = 0
							attack4tr = 0
							montr = 0
							check = 1
							hp2 = 210
							hp2or = hp2
							attacken1 = 40
							attacken2 = 80
							attacken3 = 30
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
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 3 {
											hp2 = hp2 - attack4 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
											hp2 = hp2 - attack4
											fmt.Println("Enemy took",attack4,"damage")
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
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
								fmt.Println()
								fmt.Println("choose your attack")
								fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
								fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
								fmt.Println("3.heal",attack3,"health 100% accuracy")
								fmt.Scanln(&a)
								if a == 1 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp2 = hp2 - attack1 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									}else {
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack1,"damage")
										hp2 = hp2-attack1
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
									}
								} else if a == 2 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(3)
									if random == 1 || random == 2 || random == 3 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 3 {
											hp2 = hp2 - attack2 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
										time.Sleep(1 * time.Second)
										hp2 = hp2 - attack2
										fmt.Println("Enemy took",attack2,"damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
										}
									} else {
										time.Sleep(1 * time.Second)
										fmt.Println("attack missed")
										time.Sleep(1 * time.Second)
									}
								} else if a == 3 {
									if hp1 == hp1or {
										time.Sleep(1 * time.Second)
										fmt.Println("You're not low enough to heal")
										time.Sleep(1 * time.Second)
									} else if hp1 >= (hp1or - attack3){
										time.Sleep(1 * time.Second)
										hp1 = hp1or
										fmt.Println("You have been healed to max")
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
								time.Sleep(1 * time.Second)
								fmt.Println("enemy is attacking")
								time.Sleep(1 * time.Second)
								if hp2 <= (hp2or / 2) && hp1 > attacken2 {
									time.Sleep(1 * time.Second)
									hp2 = hp2 + attacken3
									fmt.Println("enemy has been healed",attacken3,"health")
									time.Sleep(1 * time.Second)
									continue
								}
								if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp1 = hp1 - attacken1 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											check = 2
											break
										}	
									}else {
										fmt.Println("you took",attacken1,"damage")
										hp1 = hp1-attacken1
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											check = 2
											break
										}
									}
								} else {
									var b = time.Now().UnixNano()
									rand.Seed(b)
									var random3 = rand.Intn(5)
										if random3 == 1 || random3 == 2 || random3 == 3 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp1 = hp1 - attacken2 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2
													break
												}	
											}else {
												fmt.Println("you took",attacken2,"damage")
												hp1 = hp1-attacken2
												time.Sleep(1 * time.Second)	
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2
													break
												}
											}
										} else {
											fmt.Println("attack missed")
												time.Sleep(1 * time.Second)	
												continue
										}
									}
								}
								if check == 2 {
									time.Sleep(2 * time.Second)
									fmt.Println("Respawning")
									time.Sleep(5 * time.Second)
									continue
								}
							time.Sleep(4 * time.Second)
							fmt.Println("You have defeated the Skeletal Giant. you head over to the second bell, and ring it.")
							time.Sleep(4 * time.Second)
							fmt.Println("You see his hammer on the ground do you pick it up.")
							fmt.Println("1.Pick it up.")
							fmt.Println("2.Ignore it")
							fmt.Scanln(&b)
							if b == 1 {
								attack1tr = attack1
								attack2tr = attack2
								attack4tr = attack4
								hp1 = hp1or
								attack1 = 60000000
								attack2 = 80000000
								attack4 = 1200000000
								attack1tr = attack1 - attack1tr
								attack2tr = attack2 - attack2tr
								attack4tr = attack4 - attack4tr
								time.Sleep(4 * time.Second)
								fmt.Println("Weapons upgraded")
								time.Sleep(4 * time.Second)
							}
							fmt.Println("You can now advance to Mount Olympus. The largest mountain, overlooking Mystic Woods.")
							time.Sleep(4 * time.Second)
							fmt.Println("You begin an almost endless trek, and when you are about halfway there, you encounter a possessed spartan. ")
							time.Sleep(4 * time.Second)
							hp2 = 140
							hp2or = hp2
							attacken1 = 30
							attacken2 = 50
							attacken3 = 30
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
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 3 {
											hp2 = hp2 - attack4 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
											hp2 = hp2 - attack4
											fmt.Println("Enemy took",attack4,"damage")
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
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
								fmt.Println()
								fmt.Println("choose your attack")
								fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
								fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
								fmt.Println("3.heal",attack3,"health 100% accuracy")
								fmt.Scanln(&a)
								if a == 1 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp2 = hp2 - attack1 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									}else {
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack1,"damage")
										hp2 = hp2-attack1
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
									}
								} else if a == 2 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(3)
									if random == 1 || random == 2 || random == 3 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 3 {
											hp2 = hp2 - attack2 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
										time.Sleep(1 * time.Second)
										hp2 = hp2 - attack2
										fmt.Println("Enemy took",attack2,"damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
										}
									} else {
										time.Sleep(1 * time.Second)
										fmt.Println("attack missed")
										time.Sleep(1 * time.Second)
									}
								} else if a == 3 {
									if hp1 == hp1or {
										time.Sleep(1 * time.Second)
										fmt.Println("You're not low enough to heal")
										time.Sleep(1 * time.Second)
									} else if hp1 >= (hp1or - attack3){
										time.Sleep(1 * time.Second)
										hp1 = hp1or
										fmt.Println("You have been healed to max")
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
								time.Sleep(1 * time.Second)
								fmt.Println("enemy is attacking")
								time.Sleep(1 * time.Second)
								if hp2 <= (hp2or / 2) && hp1 > attacken2 {
									time.Sleep(1 * time.Second)
									hp2 = hp2 + attacken3
									fmt.Println("enemy has been healed",attacken3,"health")
									time.Sleep(1 * time.Second)
									continue
								}
								if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp1 = hp1 - attacken1 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											check = 2
											break
										}	
									}else {
										fmt.Println("you took",attacken1,"damage")
										hp1 = hp1-attacken1
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											check = 2
											break
										}
									}
								} else {
									var b = time.Now().UnixNano()
									rand.Seed(b)
									var random3 = rand.Intn(5)
										if random3 == 1 || random3 == 2 || random3 == 3 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp1 = hp1 - attacken2 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2
													break
												}	
											}else {
												fmt.Println("you took",attacken2,"damage")
												hp1 = hp1-attacken2
												time.Sleep(1 * time.Second)	
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2
													break
												}
											}
										} else {
											fmt.Println("attack missed")
												time.Sleep(1 * time.Second)	
												continue
										}
									}
								}
								if check == 2 {
									time.Sleep(2 * time.Second)
									fmt.Println("Respawning")
									time.Sleep(5 * time.Second)
									continue
								}
							hp1 = hp1or
							attack1tr = 0
							hp1tr = 0
							attack2tr = 0
							attack3tr = 0
							attack4tr = 0
							montr = 0
							break
						}
							time.Sleep(4 * time.Second)
							fmt.Println("After beating the spartan, you continue up the mountain, and see a temple at the peak.")
							time.Sleep(4 * time.Second)
							fmt.Println("You realize that is where the guardian of the gem will be, and realize you might have to verse Zeus,")
							time.Sleep(4 * time.Second)
							fmt.Println("but you are comforted realizing Zeus is an ally.")
							time.Sleep(4 * time.Second)
							fmt.Println("With this assurance, you enter the temple.You are astonished to the sight of Zeus, but you notice he looks off.")
							time.Sleep(4 * time.Second)
							fmt.Println("He has succumbed to the Curse.")
							time.Sleep(4 * time.Second)
							for {
								hp1 = hp1 - hp1tr 
								hp1 = hp1or
								attack1 = attack1 - attack1tr
								mon = mon - montr
								hp1 = hp1 - hp1tr
								hp1or = hp1
								attack2 = attack2 - attack2tr
								attack3 = attack3 - attack3tr
								attack4 = attack4 - attack4tr
								hp2or = hp2
								attack1tr = 0
								hp1tr = 0
								attack2tr = 0
								attack3tr = 0
								attack4tr = 0
								montr = 0
								check = 1
								hp2 = 300
								hp2or = hp2
								attacken1 = 60
								attacken2 = 90
								attacken3 = 30
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
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack4 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
												hp2 = hp2 - attack4
												fmt.Println("Enemy took",attack4,"damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
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
									fmt.Println()
									fmt.Println("choose your attack")
									fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
									fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
									fmt.Println("3.heal",attack3,"health 100% accuracy")
									fmt.Scanln(&a)
									if a == 1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp2 = hp2 - attack1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack1,"damage")
											hp2 = hp2-attack1
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
										}
									} else if a == 2 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(3)
										if random == 1 || random == 2 || random == 3 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 3 {
												hp2 = hp2 - attack2 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
											time.Sleep(1 * time.Second)
											hp2 = hp2 - attack2
											fmt.Println("Enemy took",attack2,"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
											}
										} else {
											time.Sleep(1 * time.Second)
											fmt.Println("attack missed")
											time.Sleep(1 * time.Second)
										}
									} else if a == 3 {
										if hp1 == hp1or {
											time.Sleep(1 * time.Second)
											fmt.Println("You're not low enough to heal")
											time.Sleep(1 * time.Second)
										} else if hp1 >= (hp1or - attack3){
											time.Sleep(1 * time.Second)
											hp1 = hp1or
											fmt.Println("You have been healed to max")
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
									time.Sleep(1 * time.Second)
									fmt.Println("enemy is attacking")
									time.Sleep(1 * time.Second)
									if hp2 <= (hp2or / 2) && hp1 > attacken2 {
										time.Sleep(1 * time.Second)
										hp2 = hp2 + attacken3
										fmt.Println("enemy has been healed",attacken3,"health")
										time.Sleep(1 * time.Second)
										continue
									}
									if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp1 = hp1 - attacken1 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}	
										}else {
											fmt.Println("you took",attacken1,"damage")
											hp1 = hp1-attacken1
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												break
											}
										}
									} else {
										var b = time.Now().UnixNano()
										rand.Seed(b)
										var random3 = rand.Intn(5)
											if random3 == 1 || random3 == 2 || random3 == 3 {
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 0 {
													hp1 = hp1 - attacken2 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
													time.Sleep(1 * time.Second)
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2
														break
													}	
												}else {
													fmt.Println("you took",attacken2,"damage")
													hp1 = hp1-attacken2
													time.Sleep(1 * time.Second)	
													if hp1 <= 0 {
														fmt.Println("You Died")
														check = 2
														break
													}
												}
											} else {
												fmt.Println("attack missed")
													time.Sleep(1 * time.Second)	
													continue
											}
										}
									}
									if check == 2 {
										time.Sleep(2 * time.Second)
										fmt.Println("Respawning")
										time.Sleep(5 * time.Second)
										continue
									}
								hp1 = hp1or
								attack1tr = 0
								hp1tr = 0
								attack2tr = 0
								attack3tr = 0
								attack4tr = 0
								montr = 0
								break
							}
								attack3 = attack3 + 15
								time.Sleep(1 * time.Second)
								fmt.Println("You Don't kill Zeus, you have knocked the curse out of him. He recovers and thanks you.")
								time.Sleep(4 * time.Second)
								fmt.Println("He gives you some of his magic")
								time.Sleep(4 * time.Second)
								fmt.Println("Magic Upgraded")
								time.Sleep(4 * time.Second)
								fmt.Println("You take the gem, and proceed to the Cursed Castle.")
								time.Sleep(7 * time.Second)
								for space := 1 ; space <=100 ; space++{
									fmt.Println()
								}
								fmt.Println("MAP 3: Cursed Castle")
								for space := 1 ; space <=20 ; space++{
									fmt.Println()
								}
								time.Sleep(2 * time.Second)	
								for space := 1 ; space <=100 ; space++{
									fmt.Println()
								}
								time.Sleep(5 * time.Second)
								fmt.Println("The Castle Gate creaks open, as you take confident steps, entering the castle.")
								time.Sleep(5 * time.Second)
								fmt.Println("The castle is dark, and cursed. A Royal Servant stops you in your path, in the path leading to the main gates.")
								time.Sleep(5 * time.Second)
								fmt.Println("How dare you enter in this blissful castle. Your presence brings shame.")
								time.Sleep(4 * time.Second)
								for {
									hp1 = hp1 - hp1tr 
									hp1 = hp1or
									attack1 = attack1 - attack1tr
									mon = mon - montr
									hp1 = hp1 - hp1tr
									hp1or = hp1
									attack2 = attack2 - attack2tr
									attack3 = attack3 - attack3tr
									attack4 = attack4 - attack4tr
									hp2or = hp2
									attack1tr = 0
									hp1tr = 0
									attack2tr = 0
									attack3tr = 0
									attack4tr = 0
									montr = 0
									check = 1
									hp2 = 200
									hp2or = hp2
									attacken1 = 40
									attacken2 = 70
									attacken3 = 30
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
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 3 {
													hp2 = hp2 - attack4 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}else {
													hp2 = hp2 - attack4
													fmt.Println("Enemy took",attack4,"damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}
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
										fmt.Println()
										fmt.Println("choose your attack")
										fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
										fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
										fmt.Println("3.heal",attack3,"health 100% accuracy")
										fmt.Scanln(&a)
										if a == 1 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp2 = hp2 - attack1 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack1,"damage")
												hp2 = hp2-attack1
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
											}
										} else if a == 2 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(3)
											if random == 1 || random == 2 || random == 3 {
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 3 {
													hp2 = hp2 - attack2 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
													time.Sleep(1 * time.Second)
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}else {
												time.Sleep(1 * time.Second)
												hp2 = hp2 - attack2
												fmt.Println("Enemy took",attack2,"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
												}
											} else {
												time.Sleep(1 * time.Second)
												fmt.Println("attack missed")
												time.Sleep(1 * time.Second)
											}
										} else if a == 3 {
											if hp1 == hp1or {
												time.Sleep(1 * time.Second)
												fmt.Println("You're not low enough to heal")
												time.Sleep(1 * time.Second)
											} else if hp1 >= (hp1or - attack3){
												time.Sleep(1 * time.Second)
												hp1 = hp1or
												fmt.Println("You have been healed to max")
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
										time.Sleep(1 * time.Second)
										fmt.Println("enemy is attacking")
										time.Sleep(1 * time.Second)
										if hp2 <= (hp2or / 2) && hp1 > attacken2 {
											time.Sleep(1 * time.Second)
											hp2 = hp2 + attacken3
											fmt.Println("enemy has been healed",attacken3,"health")
											time.Sleep(1 * time.Second)
											continue
										}
										if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp1 = hp1 - attacken1 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2
													break
												}	
											}else {
												fmt.Println("you took",attacken1,"damage")
												hp1 = hp1-attacken1
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2
													break
												}
											}
										} else {
											var b = time.Now().UnixNano()
											rand.Seed(b)
											var random3 = rand.Intn(5)
												if random3 == 1 || random3 == 2 || random3 == 3 {
													var s = time.Now().UnixNano()
													rand.Seed(s)
													var random = rand.Intn(8)
													if random == 0 {
														hp1 = hp1 - attacken2 * 3 / 2
														fmt.Println("Critical Hit")
														time.Sleep(1 * time.Second)
														fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
														time.Sleep(1 * time.Second)
														if hp1 <= 0 {
															fmt.Println("You Died")
															check = 2
															break
														}	
													}else {
														fmt.Println("you took",attacken2,"damage")
														hp1 = hp1-attacken2
														time.Sleep(1 * time.Second)	
														if hp1 <= 0 {
															fmt.Println("You Died")
															check = 2
															break
														}
													}
												} else {
													fmt.Println("attack missed")
														time.Sleep(1 * time.Second)	
														continue
												}
											}
										}
										if check == 2 {
											time.Sleep(2 * time.Second)
											fmt.Println("Respawning")
											time.Sleep(5 * time.Second)
											continue
										}
									hp1 = hp1or
									hp1 = hp1or
									attack1tr = 0
									hp1tr = 0
									attack2tr = 0
									attack3tr = 0
									attack4tr = 0
									montr = 0
									break
								}
								mon = mon + 30
								fmt.Println("You gained 30 coins")
								time.Sleep(4 * time.Second)
								fmt.Println("Total coins:",mon)
								time.Sleep(4 * time.Second)
								fmt.Println("He yells Die Cursed Scum as he dies. He was wielding a steel chest plate.")
								time.Sleep(4 * time.Second)
								fmt.Println("You claim it")
								time.Sleep(4 * time.Second)
								fmt.Println("Armour upgraded")
								time.Sleep(4 * time.Second)
								fmt.Println("You carry on through the tight corridors, and encounter a sad weapon smith.")
								time.Sleep(5 * time.Second)
								fmt.Println("He is creating a sword, longer than you have ever seen. He has a pile of a variety of weapons. you approach him.")
								time.Sleep(5 * time.Second)
								fmt.Println("Greatings ashen one, He says while looking up from his anvil")
								time.Sleep(5 * time.Second)
								fmt.Println("He hands you a purple potion")
								time.Sleep(5 * time.Second)
								fmt.Println("You will need this for your journey ahead")
								time.Sleep(5 * time.Second)
								fmt.Println("You thank him and continue forwards")
								time.Sleep(5 * time.Second)
								var poip int
								var cursep int
								var hpst int
								var rpoip = 0
								var rcursep = 0
								poip = poip + 1
								fmt.Println("Poison potion collected")
								time.Sleep(5 * time.Second)
								fmt.Println("As you continue forwards you realize spread throughout the entire castle is a mercantile system.")
								time.Sleep(5 * time.Second)
								for {
									time.Sleep(2 * time.Second)
									fmt.Println("Which shop would you like to go?")
									fmt.Println("Total Coins:",mon)
									fmt.Println("1.Weapon Shop")
									fmt.Println("2.Armor Shop")
									fmt.Println("3.Magic shop")
									fmt.Println("4.Dark potions shop")
									fmt.Println("5.Exit")
									fmt.Scanln(&a)
									time.Sleep(2 * time.Second)
									if a == 1 {
										fmt.Println("What would you like to buy?")
										fmt.Println("Total Coins:",mon)
										fmt.Println("1.Light upgrade: 30 coins")
										fmt.Println("2.Heavy upgrade: 30 coins")
										fmt.Println("3.Ultimate upgrade: 30 coins")
										fmt.Println("4.Go back")
										fmt.Scanln(&b)
										time.Sleep(2 * time.Second)
										if b == 1 {
											if mon < 30 {
												time.Sleep(1 * time.Second)
												fmt.Println("You do not have enough to buy this item")
												time.Sleep(1 * time.Second)
											}else{
												mon = mon - 30
												fmt.Println("Light attack upgraded")
												attack1 = attack1 + 10
											}
										} else if b == 2 {
											if mon < 30 {
												time.Sleep(1 * time.Second)
												fmt.Println("You do not have enough to buy this item")
												time.Sleep(1 * time.Second)
											} else {
												mon = mon - 30
												fmt.Println("Heavy attack upgraded")
												attack2 = attack2 + 10
											}
										} else if b == 3 {
											if mon < 30 {
												time.Sleep(1 * time.Second)
												fmt.Println("You do not have enough to buy this item")
												time.Sleep(1 * time.Second)
											} else {
												mon = mon - 30
												fmt.Println("Ultimate attack upgraded")
												attack4 = attack4 + 15
											}
										} else {
											continue
										}
									} else if a == 2 {
										fmt.Println("What would you like to buy?")
										fmt.Println("Total Coins:",mon)
										fmt.Println("1.Helmet upgrade: 30 coins")
										fmt.Println("2.Chestplate upgrade: 40 coins")
										fmt.Println("3.Go back")
										fmt.Scanln(&b)
										time.Sleep(2 * time.Second)
										if b == 1 {
											if mon < 30 {
												time.Sleep(1 * time.Second)
												fmt.Println("You do not have enough to buy this item")
												time.Sleep(1 * time.Second)
											} else {
												mon = mon - 30
												fmt.Println("Helmet upgraded")
												hp1 = hp1 + 10
												hp1or = hp1
											}
										} else if b == 2 {
											if mon < 40 {
												time.Sleep(1 * time.Second)
												fmt.Println("You do not have enough to buy this item")
												time.Sleep(1 * time.Second)
											} else {
												mon = mon - 40
												fmt.Println("Chestplate upgraded")
												hp1 = hp1 + 20
												hp1or = hp1
											}
										} else {
											continue
										}									
									} else if a == 3 {
										fmt.Println("What would you like to buy?")
										fmt.Println("Total Coins:",mon)
										fmt.Println("1.Small magic upgrade: 30 coins")
										fmt.Println("2.Medium magic upgrade: 40 coins")
										fmt.Println("3.Large magic upgrade: 50 coins")
										fmt.Println("4.Go back")
										fmt.Scanln(&b)
										time.Sleep(2 * time.Second)
										if b == 1 {
											if mon < 30 {
												time.Sleep(1 * time.Second)
												fmt.Println("You do not have enough to buy this item")
												time.Sleep(1 * time.Second)
											} else {
												mon = mon - 30
												fmt.Println("Magic upgraded")
												attack3 = attack3 + 10
											}
										} else if b == 2 {
											if mon < 40 {
												time.Sleep(1 * time.Second)
												fmt.Println("You do not have enough to buy this item")
												time.Sleep(1 * time.Second)
											} else {
												mon = mon - 40
												fmt.Println("Magic upgraded")
												attack3 = attack3 + 15
											}
										} else if b == 3 {
											if mon < 50 {
												time.Sleep(1 * time.Second)
												fmt.Println("You do not have enough to buy this item")
												time.Sleep(1 * time.Second)
											} else {
												mon = mon - 50
												fmt.Println("Magic upgraded")
												attack3 = attack3 + 20
											}
										} else {
											continue
										}
									} else if a == 4{
										fmt.Println("What would you like to buy?")
										fmt.Println("Total Coins:",mon)
										fmt.Println("1.Poison potion: 20 coins")
										fmt.Println("2.Curse potion: 30 coins")
										fmt.Println("3.Life steal potion: 40 coins")
										fmt.Println("4.Go back")
										fmt.Scanln(&b)
										time.Sleep(2 * time.Second)
										if b == 1 {
											if mon < 20 {
												time.Sleep(1 * time.Second)
												fmt.Println("You do not have enough to buy this item")
												time.Sleep(1 * time.Second)
											} else {
												mon = mon - 20
												fmt.Println("Dark potion added")
												poip = poip + 1
											}
										} else if b == 2 {
											if mon < 30 {
												time.Sleep(1 * time.Second)
												fmt.Println("You do not have enough to buy this item")
												time.Sleep(1 * time.Second)
											} else {
												mon = mon - 30
												fmt.Println("Dark potion added")
												cursep = cursep + 1
											}
										} else if b == 3 {
											if mon < 40 {
												time.Sleep(1 * time.Second)
												fmt.Println("You do not have enough to buy this item")
												time.Sleep(1 * time.Second)
											} else {
												mon = mon - 40
												fmt.Println("Dark potion added")
												hpst = hpst + 1
											}
										} else {
											continue
										}
									} else {
										time.Sleep(2 * time.Second)
										break
									}
								}
								for {
									check = 1
									fmt.Println("You buy what you need, with the gold you obtained from all previous battles.")
									time.Sleep(5 * time.Second)
									fmt.Println("The people of what was once known as Toadstool castle are being oppressed by the fallen king, sent by Hell Lord Hades to manage the castle.")
									time.Sleep(5 * time.Second)
									fmt.Println("The old king was beheaded in the main square of the castle. You must avenge him.")
									time.Sleep(5 * time.Second)
									fmt.Println("You continue on through the courtyard, until you realize that the castle is completely deserted of enemies. Where do you go?")
									fmt.Println("1.Throne room")
									fmt.Println("2.Exit Castle")
									fmt.Println("3.Enter Dungeon")
									fmt.Println("4.Hide")
									fmt.Scanln(&b)
									if b == 1 {
										fmt.Println("The Throne room is completely trapped.")
										time.Sleep(5 * time.Second)
										fmt.Println("You turn around to attempt to open the door, but you see the shopkeeper from Map 1. At first thought.")
										time.Sleep(5 * time.Second)
										fmt.Println("You think he is cursed, but he looks just like before. He pulls out a fiery sword. Long Live Hades, He utters")
										time.Sleep(5 * time.Second)
										fmt.Println("He pushes you to the ground and points the sword towards your throat.")
										time.Sleep(5 * time.Second)
										fmt.Println("At this moment, you realize that Mystic Woods was a scapegoat, to lure you into a trap.")
										time.Sleep(5 * time.Second)
										fmt.Println("The shopkeeper stabs the flaming sword into your back, you feel a wave of pain until you fall into a forever darkness.")
										time.Sleep(3 * time.Second)
										check = 2
									} else if b == 2 {
										fmt.Println("You exit the castle, but enemies pour out of the bushes and kill you.")
										check = 2
									} else if b == 3 {
										fmt.Println("You enter the dungeons, and see the old king, he is in bad shape.")
										time.Sleep(5 * time.Second)
										fmt.Println("He speaks to you about the Fallen King's weakness, and you now feel confident in beating him.")
										time.Sleep(5 * time.Second)
										fmt.Println("Light Attacks. You thank the King, and vow to free him when the threat is over. You proceed up the stairs.")
										time.Sleep(3 * time.Second)
									} else {
										fmt.Println("You hide but you are found by guards and executed.")
										check = 2
									}
									if check == 2 {
										time.Sleep(2 * time.Second)
										fmt.Println("Respawning")
										time.Sleep(5 * time.Second)
										continue
									}
									break
								}
								time.Sleep(3 * time.Second)
								fmt.Println("After proceeding up the stairs, you enter the Fallen kings lair, and see a terrible dream. Betrayal.")
								time.Sleep(5 * time.Second)
								fmt.Println("You ignore it and carry on. You search to fight the fallen king, but encounter A Royal Servant!")
								for {
									hp1 = hp1 - hp1tr 
									hp1 = hp1or
									attack1 = attack1 - attack1tr
									mon = mon - montr
									hp1 = hp1 - hp1tr
									hp1or = hp1
									attack2 = attack2 - attack2tr
									attack3 = attack3 - attack3tr
									attack4 = attack4 - attack4tr
									hp2or = hp2
									attack1tr = 0
									hp1tr = 0
									attack2tr = 0
									attack3tr = 0
									attack4tr = 0
									montr = 0
									check = 1
									rpoip = 0
									rcursep = 0
									hp2 = 350
									hp2or = hp2
									attacken1 = 50
									attacken2 = 80
									attacken3 = 30
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
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 3 {
													hp2 = hp2 - attack4 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}else {
													hp2 = hp2 - attack4
													fmt.Println("Enemy took",attack4,"damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}
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
										fmt.Println()
										fmt.Println("choose your attack")
										fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
										fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
										fmt.Println("3.heal",attack3,"health 100% accuracy")
										fmt.Println("4.dark magic 100% accuracy")
										fmt.Scanln(&a)
										if a == 1 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp2 = hp2 - attack1 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack1,"damage")
												hp2 = hp2-attack1
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
											}
										} else if a == 2 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(3)
											if random == 1 || random == 2 || random == 3 {
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 3 {
													hp2 = hp2 - attack2 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
													time.Sleep(1 * time.Second)
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}else {
												time.Sleep(1 * time.Second)
												hp2 = hp2 - attack2
												fmt.Println("Enemy took",attack2,"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
												}
											} else {
												time.Sleep(1 * time.Second)
												fmt.Println("attack missed")
												time.Sleep(1 * time.Second)
											}
										} else if a == 3 {
											if hp1 == hp1or {
												time.Sleep(1 * time.Second)
												fmt.Println("You're not low enough to heal")
												time.Sleep(1 * time.Second)
											} else if hp1 >= (hp1or - attack3){
												time.Sleep(1 * time.Second)
												hp1 = hp1or
												fmt.Println("You have been healed to max")
												time.Sleep(1 * time.Second)
											} else {
												time.Sleep(1 * time.Second)
												hp1 = hp1 + attack3
												fmt.Println("You have been healed",attack3,"health")
												time.Sleep(1 * time.Second)
											}
										} else if a == 4 {
											time.Sleep(1 * time.Second)
											fmt.Println("Which potion would you like to use")
											fmt.Println("1.Poison:",poip)
											fmt.Println("2.Curse",cursep)
											fmt.Println("3.Life Steal",hpst)
											fmt.Scanln(&b)
											time.Sleep(1 * time.Second)
											if b == 1 {
												if poip == 0 {
													time.Sleep(1 * time.Second)
													fmt.Println("Sorry you don't have this item")
													time.Sleep(1 * time.Second)
													continue
												} else {
													poip--
													time.Sleep(1 * time.Second)
													rpoip = 5
												}
											} else if b == 2 {
												if cursep == 0 {
													time.Sleep(1 * time.Second)
													fmt.Println("Sorry you don't have this item")
													time.Sleep(1 * time.Second)
													continue
												} else {
													cursep--
													time.Sleep(1 * time.Second)
													rcursep = 10
												}
											} else if b == 3 {
												if hpst == 0 {
													time.Sleep(1 * time.Second)
													fmt.Println("Sorry you don't have this item")
													time.Sleep(1 * time.Second)
													continue
												} else {
													hpst--
													if hp1 == hp1or {
														hp2 = hp2 - 30
														time.Sleep(1 * time.Second)
														fmt.Println("You're not low enough to heal")
														time.Sleep(1 * time.Second)
														fmt.Println("Enemy took 30 damage")
													} else if hp1 >= (hp1or - 30){
														time.Sleep(1 * time.Second)
														hp1 = hp1or
														hp2 = hp2 - 30
														fmt.Println("You have been healed to max")
														time.Sleep(1 * time.Second)
														fmt.Println("Enemy took 30 damage")
														if hp2 <= 0 {
															fmt.Println("You Win")
															break
														}
													} else {
														time.Sleep(1 * time.Second)
														hp1 = hp1 + 30
														hp2 = hp2 - 30
														fmt.Println("You stole 30 hp")
														if hp2 <= 0 {
															fmt.Println("You Win")
															break
														}
													}
												}
											}
										} else {
											fmt.Println("Sorry that is an invalid command")
										}
										time.Sleep(1 * time.Second)
										fmt.Println("enemy is attacking")
										time.Sleep(1 * time.Second)
										if hp2 <= (hp2or / 2) && hp1 > attacken2 {
											time.Sleep(1 * time.Second)
											hp2 = hp2 + attacken3
											fmt.Println("enemy has been healed",attacken3,"health")
											time.Sleep(1 * time.Second)
											continue
										}
										if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp1 = hp1 - attacken1 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2 
													break
												}	
											}else {
												fmt.Println("you took",attacken1,"damage")
												hp1 = hp1-attacken1
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2 
													break
												}
											}
										} else {
											var b = time.Now().UnixNano()
											rand.Seed(b)
											var random3 = rand.Intn(5)
												if random3 == 1 || random3 == 2 || random3 == 3 {
													var s = time.Now().UnixNano()
													rand.Seed(s)
													var random = rand.Intn(8)
													if random == 0 {
														hp1 = hp1 - attacken2 * 3 / 2
														fmt.Println("Critical Hit")
														time.Sleep(1 * time.Second)
														fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
														time.Sleep(1 * time.Second)
														if hp1 <= 0 {
															fmt.Println("You Died")
															check = 2 
															break
														}	
													}else {
														fmt.Println("you took",attacken2,"damage")
														hp1 = hp1-attacken2
														time.Sleep(1 * time.Second)	
														if hp1 <= 0 {
															fmt.Println("You Died")
															check = 2 
															break
														}
													}
												} else {
													fmt.Println("attack missed")
														time.Sleep(1 * time.Second)	
														continue
												}
											}
											if rpoip != 0 {
												hp2 = hp2 - 10
												rpoip--
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy poisoned 10 damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}
											if rcursep != 0 {
												if rcursep == 1 {
													hp2 = hp2 - 150
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took 150 damage.")
													time.Sleep(1 * time.Second)
													rcursep = 0
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}
												rcursep--
												time.Sleep(1 * time.Second)
												fmt.Println(rcursep,"turns left")
												time.Sleep(1 * time.Second)
											}
										}
									hp1 = hp1or
									if check == 2 {
										time.Sleep(2 * time.Second)
										fmt.Println("Respawning")
										time.Sleep(5 * time.Second)
										continue
									}
								hp1 = hp1or
								hp1 = hp1or
								attack1tr = 0
								hp1tr = 0
								attack2tr = 0
								attack3tr = 0
								attack4tr = 0
								montr = 0
								break
								}
								time.Sleep(3 * time.Second)
								fmt.Println("After beating the servant, you realize he was different from the others.")
								time.Sleep(5 * time.Second)
								fmt.Println("He is one of the Fallen King's right-hand men.")
								time.Sleep(5 * time.Second)
								fmt.Println("He was leaving to kill the King in the dungeons.")
								time.Sleep(5 * time.Second)
								fmt.Println("Wondering where the Fallen King was, you exit the throne room, and see Terror, Death, and Blood.")
								time.Sleep(5 * time.Second)
								fmt.Println("The Merchants were killed, impaled on stakes for treachery. Enraged, you go back to the Throne Room.")
								time.Sleep(5 * time.Second)
								fmt.Println("The Fallen King is there and you get ready for battle")
								for {
									hp1 = hp1 - hp1tr 
									hp1 = hp1or
									attack1 = attack1 - attack1tr
									mon = mon - montr
									hp1 = hp1 - hp1tr
									hp1or = hp1
									attack2 = attack2 - attack2tr
									attack3 = attack3 - attack3tr
									attack4 = attack4 - attack4tr
									hp2or = hp2
									attack1tr = 0
									hp1tr = 0
									attack2tr = 0
									attack3tr = 0
									attack4tr = 0
									montr = 0
									check = 1
									rpoip = 0
									rcursep = 0
									hp2 = 1100
									hp2or = hp2
									attacken1 = 50
									attacken2 = 120
									attacken3 = 30
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
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 3 {
													hp2 = hp2 - attack4 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}else {
													hp2 = hp2 - attack4
													fmt.Println("Enemy took",attack4,"damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}
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
										fmt.Println()
										fmt.Println("choose your attack")
										fmt.Println("1.light attack ",attack1 * 2,"damage 100% accuracy")
										fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
										fmt.Println("3.heal",attack3,"health 100% accuracy")
										fmt.Println("4.dark magic 100% accuracy")
										fmt.Scanln(&a)
										if a == 1 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp2 = hp2 - attack1 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack1 * 2 + (attack1),"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											} else {
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack1 * 2,"damage")
												hp2 = hp2 - attack1 * 2
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
											}
										} else if a == 2 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(3)
											if random == 1 || random == 2 || random == 3 {
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 3 {
													hp2 = hp2 - attack2 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
													time.Sleep(1 * time.Second)
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}else {
												time.Sleep(1 * time.Second)
												hp2 = hp2 - attack2
												fmt.Println("Enemy took",attack2,"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
												}
											} else {
												time.Sleep(1 * time.Second)
												fmt.Println("attack missed")
												time.Sleep(1 * time.Second)
											}
										} else if a == 3 {
											if hp1 == hp1or {
												time.Sleep(1 * time.Second)
												fmt.Println("You're not low enough to heal")
												time.Sleep(1 * time.Second)
											} else if hp1 >= (hp1or - attack3){
												time.Sleep(1 * time.Second)
												hp1 = hp1or
												fmt.Println("You have been healed to max")
												time.Sleep(1 * time.Second)
											} else {
												time.Sleep(1 * time.Second)
												hp1 = hp1 + attack3
												fmt.Println("You have been healed",attack3,"health")
												time.Sleep(1 * time.Second)
											}
										} else if a == 4 {
											time.Sleep(1 * time.Second)
											fmt.Println("Which potion would you like to use")
											fmt.Println("1.Poison:",poip)
											fmt.Println("2.Curse",cursep)
											fmt.Println("3.Life Steal",hpst)
											fmt.Scanln(&b)
											time.Sleep(1 * time.Second)
											if b == 1 {
												if poip == 0 {
													time.Sleep(1 * time.Second)
													fmt.Println("Sorry you don't have this item")
													time.Sleep(1 * time.Second)
													continue
												} else {
													poip--
													time.Sleep(1 * time.Second)
													rpoip = 5
												}
											} else if b == 2 {
												if cursep == 0 {
													time.Sleep(1 * time.Second)
													fmt.Println("Sorry you don't have this item")
													time.Sleep(1 * time.Second)
													continue
												} else {
													cursep--
													time.Sleep(1 * time.Second)
													rcursep = 10
												}
											} else if b == 3 {
												if hpst == 0 {
													time.Sleep(1 * time.Second)
													fmt.Println("Sorry you don't have this item")
													time.Sleep(1 * time.Second)
													continue
												} else {
													hpst--
													if hp1 == hp1or {
														hp2 = hp2 - 30
														time.Sleep(1 * time.Second)
														fmt.Println("You're not low enough to heal")
														time.Sleep(1 * time.Second)
														fmt.Println("Enemy took 30 damage")
													} else if hp1 >= (hp1or - 30){
														time.Sleep(1 * time.Second)
														hp1 = hp1or
														hp2 = hp2 - 30
														fmt.Println("You have been healed to max")
														time.Sleep(1 * time.Second)
														fmt.Println("Enemy took 30 damage")
														if hp2 <= 0 {
															fmt.Println("You Win")
															break
														}
													} else {
														time.Sleep(1 * time.Second)
														hp1 = hp1 + 30
														hp2 = hp2 - 30
														fmt.Println("You stole 30 hp")
														if hp2 <= 0 {
															fmt.Println("You Win")
															break
														}
													}
												}
											}
										} else {
											fmt.Println("Sorry that is an invalid command")
										}
										time.Sleep(1 * time.Second)
										fmt.Println("enemy is attacking")
										time.Sleep(1 * time.Second)
										if hp2 <= 300 && hp1 > attacken2 {
											time.Sleep(1 * time.Second)
											hp2 = hp2 + attacken3
											fmt.Println("enemy has been healed",attacken3,"health")
											time.Sleep(1 * time.Second)
											continue
										}
										if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp1 = hp1 - attacken1 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2 
													break
												}	
											} else {
												fmt.Println("you took",attacken1,"damage")
												hp1 = hp1-attacken1
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2 
													break
												}
											}
										} else {
											var b = time.Now().UnixNano()
											rand.Seed(b)
											var random3 = rand.Intn(5)
												if random3 == 1 || random3 == 2 || random3 == 3 {
													var s = time.Now().UnixNano()
													rand.Seed(s)
													var random = rand.Intn(8)
													if random == 0 {
														hp1 = hp1 - attacken2 * 3 / 2
														fmt.Println("Critical Hit")
														time.Sleep(1 * time.Second)
														fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
														time.Sleep(1 * time.Second)
														if hp1 <= 0 {
															fmt.Println("You Died")
															check = 2 
															break
														}	
													}else {
														fmt.Println("you took",attacken2,"damage")
														hp1 = hp1-attacken2
														time.Sleep(1 * time.Second)	
														if hp1 <= 0 {
															fmt.Println("You Died")
															check = 2 
															break
														}
													}
												} else {
													fmt.Println("attack missed")
														time.Sleep(1 * time.Second)	
														continue
												}
											}
											if rpoip != 0 {
												hp2 = hp2 - 10
												rpoip--
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy poisoned 10 damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}
											if rcursep != 0 {
												if rcursep == 1 {
													hp2 = hp2 - 150
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took 150 damage.")
													time.Sleep(1 * time.Second)
													rcursep = 0
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}
												rcursep--
												time.Sleep(1 * time.Second)
												fmt.Println(rcursep,"turns left")
												time.Sleep(1 * time.Second)
											}
										}
									hp1 = hp1or
									if check == 2 {
										time.Sleep(2 * time.Second)
										fmt.Println("Respawning")
										time.Sleep(5 * time.Second)
										continue
									}
								hp1 = hp1or
								hp1 = hp1or
								attack1tr = 0
								hp1tr = 0
								attack2tr = 0
								attack3tr = 0
								attack4tr = 0
								montr = 0
								break
								}
								fmt.Println("after the battle you turn around to see Hades. He effortlessly captures you alongside his men.")
								time.Sleep(5 * time.Second)
								for space := 1 ; space <=100 ; space++{
									fmt.Println()
								}
								fmt.Println("MAP 4: HELL'S GATE")
								for space := 1 ; space <=20 ; space++{
									fmt.Println()
								}
								time.Sleep(2 * time.Second)	
								for space := 1 ; space <=100 ; space++{
									fmt.Println()
								}
								time.Sleep(2 * time.Second)
								fmt.Println("You wake up the the endless bumping of the wheels of the caged cart rolling on the Lava terrain of Hell's Gate.")
								time.Sleep(5 * time.Second)
								fmt.Println("You are being rolled towards a castle, which is still about 2000m away. There is a sad-looking wolf in the cage.")
								time.Sleep(5 * time.Second)
								fmt.Println("Suddenly as you are about to doze off, the entire cage tips over. The wheel fell off.")
								time.Sleep(5 * time.Second)
								for { 
									fmt.Println("You see an opportunity to escape. Do you take it? ")
									fmt.Println("1.Yes")
									fmt.Println("2.No")
									fmt.Scanln(&b)
									time.Sleep(2 * time.Second)
									if b == 1 {
										fmt.Println("You force the lock open, and you and the wolf run, and the guards notice once you pick up speed.")
										time.Sleep(5 * time.Second)
										fmt.Println("The chase after you and shoot you with flame arrows, but they miss.")
										time.Sleep(5 * time.Second)
										fmt.Println("You run into Hellbent Woods, and locate a sleeping elite guard. Stealthily, you go around him and twist his neck.")
										time.Sleep(5 * time.Second)
										fmt.Println("You take his sword, and enter battle with the Guard.")
										attack2 = 130000000
										attack1 = 900000000
										attack4 = 2000000000
									} else {
										time.Sleep(2 * time.Second)
										fmt.Println("The cart gets fixed, and Hades executes you ")
										time.Sleep(2 * time.Second)
										fmt.Println("Respawning")
										time.Sleep(5 * time.Second)
										continue
									}
									break
								}
								for{
									hp1 = hp1 - hp1tr 
									hp1 = hp1or
									attack1 = attack1 - attack1tr
									mon = mon - montr
									hp1 = hp1 - hp1tr
									hp1or = hp1
									attack2 = attack2 - attack2tr
									attack3 = attack3 - attack3tr
									attack4 = attack4 - attack4tr
									hp2or = hp2
									attack1tr = 0
									hp1tr = 0
									attack2tr = 0
									attack3tr = 0
									attack4tr = 0
									montr = 0
									check = 1
									rpoip = 0
									rcursep = 0
									hp2 = 400
									hp2or = hp2
									attacken1 = 60
									attacken2 = 100
									attacken3 = 40
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
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 3 {
													hp2 = hp2 - attack4 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}else {
													hp2 = hp2 - attack4
													fmt.Println("Enemy took",attack4,"damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}
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
										fmt.Println()
										fmt.Println("choose your attack")
										fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
										fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
										fmt.Println("3.heal",attack3,"health 100% accuracy")
										fmt.Println("4.dark magic 100% accuracy")
										fmt.Scanln(&a)
										if a == 1 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp2 = hp2 - attack1 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack1,"damage")
												hp2 = hp2 - attack1
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
											}
										} else if a == 2 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(3)
											if random == 1 || random == 2 || random == 3 {
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 3 {
													hp2 = hp2 - attack2 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
													time.Sleep(1 * time.Second)
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}else {
												time.Sleep(1 * time.Second)
												hp2 = hp2 - attack2
												fmt.Println("Enemy took",attack2,"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
												}
											} else {
												time.Sleep(1 * time.Second)
												fmt.Println("attack missed")
												time.Sleep(1 * time.Second)
											}
										} else if a == 3 {
											if hp1 == hp1or {
												time.Sleep(1 * time.Second)
												fmt.Println("You're not low enough to heal")
												time.Sleep(1 * time.Second)
											} else if hp1 >= (hp1or - attack3){
												time.Sleep(1 * time.Second)
												hp1 = hp1or
												fmt.Println("You have been healed to max")
												time.Sleep(1 * time.Second)
											} else {
												time.Sleep(1 * time.Second)
												hp1 = hp1 + attack3
												fmt.Println("You have been healed",attack3,"health")
												time.Sleep(1 * time.Second)
											}
										} else if a == 4 {
											time.Sleep(1 * time.Second)
											fmt.Println("Which potion would you like to use")
											fmt.Println("1.Poison:",poip)
											fmt.Println("2.Curse",cursep)
											fmt.Println("3.Life Steal",hpst)
											fmt.Scanln(&b)
											time.Sleep(1 * time.Second)
											if b == 1 {
												if poip == 0 {
													time.Sleep(1 * time.Second)
													fmt.Println("Sorry you don't have this item")
													time.Sleep(1 * time.Second)
													continue
												} else {
													poip--
													time.Sleep(1 * time.Second)
													rpoip = 5
												}
											} else if b == 2 {
												if cursep == 0 {
													time.Sleep(1 * time.Second)
													fmt.Println("Sorry you don't have this item")
													time.Sleep(1 * time.Second)
													continue
												} else {
													cursep--
													time.Sleep(1 * time.Second)
													rcursep = 10
												}
											} else if b == 3 {
												if hpst == 0 {
													time.Sleep(1 * time.Second)
													fmt.Println("Sorry you don't have this item")
													time.Sleep(1 * time.Second)
													continue
												} else {
													hpst--
													if hp1 == hp1or {
														hp2 = hp2 - 30
														time.Sleep(1 * time.Second)
														fmt.Println("You're not low enough to heal")
														time.Sleep(1 * time.Second)
														fmt.Println("Enemy took 30 damage")
													} else if hp1 >= (hp1or - 30){
														time.Sleep(1 * time.Second)
														hp1 = hp1or
														hp2 = hp2 - 30
														fmt.Println("You have been healed to max")
														time.Sleep(1 * time.Second)
														fmt.Println("Enemy took 30 damage")
														if hp2 <= 0 {
															fmt.Println("You Win")
															break
														}
													} else {
														time.Sleep(1 * time.Second)
														hp1 = hp1 + 30
														hp2 = hp2 - 30
														fmt.Println("You stole 30 hp")
														if hp2 <= 0 {
															fmt.Println("You Win")
															break
														}
													}
												}
											}
										} else {
											fmt.Println("Sorry that is an invalid command")
										}
										time.Sleep(1 * time.Second)
										fmt.Println("enemy is attacking")
										time.Sleep(1 * time.Second)
										if hp2 <= (hp2or / 2) && hp1 > attacken2 {
											time.Sleep(1 * time.Second)
											hp2 = hp2 + attacken3
											fmt.Println("enemy has been healed",attacken3,"health")
											time.Sleep(1 * time.Second)
											continue
										}
										if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp1 = hp1 - attacken1 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2 
													break
												}	
											}else {
												fmt.Println("you took",attacken1,"damage")
												hp1 = hp1-attacken1
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2 
													break
												}
											}
										} else {
											var b = time.Now().UnixNano()
											rand.Seed(b)
											var random3 = rand.Intn(5)
												if random3 == 1 || random3 == 2 || random3 == 3 {
													var s = time.Now().UnixNano()
													rand.Seed(s)
													var random = rand.Intn(8)
													if random == 0 {
														hp1 = hp1 - attacken2 * 3 / 2
														fmt.Println("Critical Hit")
														time.Sleep(1 * time.Second)
														fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
														time.Sleep(1 * time.Second)
														if hp1 <= 0 {
															fmt.Println("You Died")
															check = 2 
															break
														}	
													}else {
														fmt.Println("you took",attacken2,"damage")
														hp1 = hp1-attacken2
														time.Sleep(1 * time.Second)	
														if hp1 <= 0 {
															fmt.Println("You Died")
															check = 2 
															break
														}
													}
												} else {
													fmt.Println("attack missed")
														time.Sleep(1 * time.Second)	
														continue
												}
											}
											if rpoip != 0 {
												hp2 = hp2 - 10
												rpoip--
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy poisoned 10 damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}
											if rcursep != 0 {
												if rcursep == 1 {
													hp2 = hp2 - 150
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took 150 damage.")
													time.Sleep(1 * time.Second)
													rcursep = 0
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}
												rcursep--
												time.Sleep(1 * time.Second)
												fmt.Println(rcursep,"turns left")
												time.Sleep(1 * time.Second)
											}
										}
									if check == 2 {
										time.Sleep(2 * time.Second)
										fmt.Println("Respawning")
										time.Sleep(5 * time.Second)
										continue
									}
									if check == 2 {
										time.Sleep(2 * time.Second)
										fmt.Println("Respawning")
										time.Sleep(5 * time.Second)
										continue
									}
								hp1 = hp1or
								attack1tr = 0
								hp1tr = 0
								attack2tr = 0
								attack3tr = 0
								attack4tr = 0
								montr = 0
								break
								}
								var petat = 50
								var pethe = 40
								var toadkiller = 0
								time.Sleep(2 * time.Second)
								fmt.Println("You continue with the wolf. The Wolf has now became your companion.")
								time.Sleep(5 * time.Second)
								fmt.Println("You reclaim your armor from the guard, and proceed towards the castle, where Hades is at the moment.")
								time.Sleep(5 * time.Second)
								fmt.Println("He expects you as a prisoner in his presence, but little does he know, he will see fate in a matter of time.")
								time.Sleep(5 * time.Second)
								fmt.Println("You run towards the castle, but feel something grab your arm. It is a Hades Council Member.")
								time.Sleep(5 * time.Second)
								fmt.Println("He says Greetings, Ashen One You realize he is a spy.")
								time.Sleep(5 * time.Second)
								fmt.Println("he takes you into the Woods, and see a system of refugees in a camp. There are shops there.")
								time.Sleep(5 * time.Second)
								var shop1 = 0
								var shop2 = 0
								var shop3 = 0
								var shop4 = 0
								var shop5 = 0
								for {
									fmt.Println("Where would you like to go?")
									fmt.Println("1.Pet shop")
									fmt.Println("2.Weapon shop")
									fmt.Println("3.Shady shop")
									fmt.Println("4.Armor shop")
									fmt.Println("5.Dark potion shop")
									fmt.Println("6.Exit")
									fmt.Scanln(&b)
									time.Sleep(1 * time.Second)
									if b == 1 {
										if shop1 == 0 {
											shop1++
											petat = petat + 10
											pethe = pethe + 10
											fmt.Println("The merchent gives your wolf some food and gives him some armor")
											time.Sleep(1 * time.Second)
										} else {
											fmt.Println("The merchent gives your wolf some food")
											time.Sleep(1 * time.Second)
										}
									} else if b == 2 {
										if shop2 == 0 {
											shop2++
											attack1 = attack1 + 10
											attack2 = attack2 + 10
											attack4 = attack4 + 20
											fmt.Println("The merchent sharpen your weapon blade")
											time.Sleep(1 * time.Second)
										} else {
											fmt.Println("The merchent inspects your weapon but it is still in good condition")
											time.Sleep(1 * time.Second)
										}
									} else if b == 3 {
										if shop3 == 0 {
											shop3++
											toadkiller = 1
											fmt.Println("The merchent hands you a strange potion Use this to beat the one with the shroom. He says")
											time.Sleep(5 * time.Second)
											fmt.Println("You look at the potion. It doesn't look like anything you've seen before. You thank the merchnt though and continue shopping.")
											time.Sleep(1 * time.Second)
											} else {
											fmt.Println("The merchent wishes you luck for the future")
											time.Sleep(1 * time.Second)
										}
									} else if b == 4 {
										if shop4 == 0 {
											shop4++
											hp1 = hp1 + 50
											hp1or = hp1
											fmt.Println("The merchent fixes your armor")
											time.Sleep(1 * time.Second)
										} else {
											fmt.Println("The merchent says hi")
											time.Sleep(1 * time.Second)
										}
									} else if b == 5 {
										if shop5 == 0 {
											shop5++
											cursep = cursep + 1
											fmt.Println("The merchent gives you a curse potion")
											time.Sleep(1 * time.Second)
										} else {
											fmt.Println("The merchent says hi")
											time.Sleep(1 * time.Second)
										}
									} else {
										time.Sleep(1 * time.Second)
										break
									}
								}
								if toadkiller == 0 {
									toadkiller = 1
									fmt.Println("As you exit a shop a merchent stops you. The merchent hands you a strange potion Use this to beat the one with the shroom. He says")
									time.Sleep(5 * time.Second)
									fmt.Println("You look at the potion. It doesn't look like anything you've seen before. You thank the merchnt though and continue shopping.")
									time.Sleep(1 * time.Second)
								}
								fmt.Println("After your business, you exit the Black Market, and are greeted by the heat of the pouring lava from the Magma.")
								time.Sleep(5 * time.Second)
								fmt.Println("You notice Hades castle, and head towards it...The castle gates open, and you are greeted by Hades' best man.")
								time.Sleep(5 * time.Second)
								fmt.Println("Halt, Who goes there! He shouts.")
								for {
									hp1 = hp1 - hp1tr 
									hp1 = hp1or
									attack1 = attack1 - attack1tr
									mon = mon - montr
									hp1 = hp1 - hp1tr
									hp1or = hp1
									attack2 = attack2 - attack2tr
									attack3 = attack3 - attack3tr
									attack4 = attack4 - attack4tr
									hp2or = hp2
									attack1tr = 0
									hp1tr = 0
									attack2tr = 0
									attack3tr = 0
									attack4tr = 0
									montr = 0
									check = 1
									rpoip = 0
									rcursep = 0
									hp2 = 600
									hp2or = hp2
									attacken1 = 70
									attacken2 = 110
									attacken3 = 40
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
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 3 {
													hp2 = hp2 - attack4 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}else {
													hp2 = hp2 - attack4
													fmt.Println("Enemy took",attack4,"damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}
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
										fmt.Println()
										fmt.Println("choose your attack")
										fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
										fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
										fmt.Println("3.heal",attack3,"health 100% accuracy")
										fmt.Println("4.dark magic 100% accuracy")
										fmt.Scanln(&a)
										if a == 1 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp2 = hp2 - attack1 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack1,"damage")
												hp2 = hp2 - attack1
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
											}
										} else if a == 2 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(3)
											if random == 1 || random == 2 || random == 3 {
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 3 {
													hp2 = hp2 - attack2 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
													time.Sleep(1 * time.Second)
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}else {
												time.Sleep(1 * time.Second)
												hp2 = hp2 - attack2
												fmt.Println("Enemy took",attack2,"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
												}
											} else {
												time.Sleep(1 * time.Second)
												fmt.Println("attack missed")
												time.Sleep(1 * time.Second)
											}
										} else if a == 3 {
											if hp1 == hp1or {
												time.Sleep(1 * time.Second)
												fmt.Println("You're not low enough to heal")
												time.Sleep(1 * time.Second)
											} else if hp1 >= (hp1or - attack3){
												time.Sleep(1 * time.Second)
												hp1 = hp1or
												fmt.Println("You have been healed to max")
												time.Sleep(1 * time.Second)
											} else {
												time.Sleep(1 * time.Second)
												hp1 = hp1 + attack3
												fmt.Println("You have been healed",attack3,"health")
												time.Sleep(1 * time.Second)
											}
										} else if a == 4 {
											time.Sleep(1 * time.Second)
											fmt.Println("Which potion would you like to use")
											fmt.Println("1.Poison:",poip)
											fmt.Println("2.Curse",cursep)
											fmt.Println("3.Life Steal",hpst)
											fmt.Scanln(&b)
											time.Sleep(1 * time.Second)
											if b == 1 {
												if poip == 0 {
													time.Sleep(1 * time.Second)
													fmt.Println("Sorry you don't have this item")
													time.Sleep(1 * time.Second)
													continue
												} else {
													poip--
													time.Sleep(1 * time.Second)
													rpoip = 5
												}
											} else if b == 2 {
												if cursep == 0 {
													time.Sleep(1 * time.Second)
													fmt.Println("Sorry you don't have this item")
													time.Sleep(1 * time.Second)
													continue
												} else {
													cursep--
													time.Sleep(1 * time.Second)
													rcursep = 10
												}
											} else if b == 3 {
												if hpst == 0 {
													time.Sleep(1 * time.Second)
													fmt.Println("Sorry you don't have this item")
													time.Sleep(1 * time.Second)
													continue
												} else {
													hpst--
													if hp1 == hp1or {
														hp2 = hp2 - 30
														time.Sleep(1 * time.Second)
														fmt.Println("You're not low enough to heal")
														time.Sleep(1 * time.Second)
														fmt.Println("Enemy took 30 damage")
													} else if hp1 >= (hp1or - 30){
														time.Sleep(1 * time.Second)
														hp1 = hp1or
														hp2 = hp2 - 30
														fmt.Println("You have been healed to max")
														time.Sleep(1 * time.Second)
														fmt.Println("Enemy took 30 damage")
														if hp2 <= 0 {
															fmt.Println("You Win")
															break
														}
													} else {
														time.Sleep(1 * time.Second)
														hp1 = hp1 + 30
														hp2 = hp2 - 30
														fmt.Println("You stole 30 hp")
														if hp2 <= 0 {
															fmt.Println("You Win")
															break
														}
													}
												}
											}
										} else {
											fmt.Println("Sorry that is an invalid command")
										}
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(5)
										if random == 3 {
											if hp1 == hp1or / 2 {
												hp1 = hp1 + pethe
												time.Sleep(1 * time.Second)
												fmt.Println("Your pet healed you",pethe,"hp")
											} else {
												hp2 = hp2 - petat
												time.Sleep(1 * time.Second)
												fmt.Println("Your pet did",petat,"damage")
											}
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}
										time.Sleep(1 * time.Second)
										fmt.Println("enemy is attacking")
										time.Sleep(1 * time.Second)
										if hp2 <= (hp2or / 2) && hp1 > attacken2 {
											time.Sleep(1 * time.Second)
											hp2 = hp2 + attacken3
											fmt.Println("enemy has been healed",attacken3,"health")
											time.Sleep(1 * time.Second)
											continue
										}
										if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp1 = hp1 - attacken1 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2 
													break
												}	
											}else {
												fmt.Println("you took",attacken1,"damage")
												hp1 = hp1-attacken1
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2 
													break
												}
											}
										} else {
											var b = time.Now().UnixNano()
											rand.Seed(b)
											var random3 = rand.Intn(5)
												if random3 == 1 || random3 == 2 || random3 == 3 {
													var s = time.Now().UnixNano()
													rand.Seed(s)
													var random = rand.Intn(8)
													if random == 0 {
														hp1 = hp1 - attacken2 * 3 / 2
														fmt.Println("Critical Hit")
														time.Sleep(1 * time.Second)
														fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
														time.Sleep(1 * time.Second)
														if hp1 <= 0 {
															fmt.Println("You Died")
															check = 2 
															break
														}	
													}else {
														fmt.Println("you took",attacken2,"damage")
														hp1 = hp1-attacken2
														time.Sleep(1 * time.Second)	
														if hp1 <= 0 {
															fmt.Println("You Died")
															check = 2 
															break
														}
													}
												} else {
													fmt.Println("attack missed")
														time.Sleep(1 * time.Second)	
														continue
												}
											}
											if rpoip != 0 {
												hp2 = hp2 - 10
												rpoip--
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy poisoned 10 damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}
											if rcursep != 0 {
												if rcursep == 1 {
													hp2 = hp2 - 150
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took 150 damage.")
													time.Sleep(1 * time.Second)
													rcursep = 0
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}
												rcursep--
												time.Sleep(1 * time.Second)
												fmt.Println(rcursep,"turns left")
												time.Sleep(1 * time.Second)
											}
										}
										if check == 2 {
											time.Sleep(2 * time.Second)
											fmt.Println("Respawning")
											time.Sleep(5 * time.Second)
											continue
										}
									hp1 = hp1or
									attack1tr = 0
									hp1tr = 0
									attack2tr = 0
									attack3tr = 0
									attack4tr = 0
									montr = 0
									break
								}
								fmt.Println("After beating the Top Soldier, you claim all his armor.")
								time.Sleep(5 * time.Second)
								fmt.Println("You can go forge your armor with your existing armor, or continue. What do you do?")
								fmt.Println("1.Forge")
								fmt.Println("2.Don't forge")
								fmt.Scanln(&b)
								if b == 1 {
									hp1 = hp1 + 60
									hp1or = hp1
									time.Sleep(1 * time.Second)
									fmt.Println("You re-enter the black market, and ask the Blacksmith for the forging. He agrees, and wishes you luck.")
									time.Sleep(1 * time.Second)
								}
								time.Sleep(1 * time.Second)
								fmt.Println("Now, you are confident, so you go back to the Castle, and head straight for Hades throne.")
								time.Sleep(5 * time.Second)
								fmt.Println("Hades is sitting there, and is surprised to see you. How did you escape!")
								time.Sleep(5 * time.Second)
								fmt.Println("You say, Its a Work of art, something a horned creep wouldn't know")
								for {
									hp1 = hp1 - hp1tr 
									hp1 = hp1or
									attack1 = attack1 - attack1tr
									mon = mon - montr
									hp1 = hp1 - hp1tr
									hp1or = hp1
									attack2 = attack2 - attack2tr
									attack3 = attack3 - attack3tr
									attack4 = attack4 - attack4tr
									hp2or = hp2
									attack1tr = 0
									hp1tr = 0
									attack2tr = 0
									attack3tr = 0
									attack4tr = 0
									montr = 0
									check = 1
									rpoip = 0
									rcursep = 0
									hp2 = 750
									hp2or = hp2
									attacken1 = 80
									attacken2 = 120
									attacken3 = 50
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
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 3 {
													hp2 = hp2 - attack4 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}else {
													hp2 = hp2 - attack4
													fmt.Println("Enemy took",attack4,"damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}
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
										fmt.Println()
										fmt.Println("choose your attack")
										fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
										fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
										fmt.Println("3.heal",attack3,"health 100% accuracy")
										fmt.Println("4.dark magic 100% accuracy")
										fmt.Scanln(&a)
										if a == 1 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp2 = hp2 - attack1 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack1,"damage")
												hp2 = hp2 - attack1
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
											}
										} else if a == 2 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(3)
											if random == 1 || random == 2 || random == 3 {
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 3 {
													hp2 = hp2 - attack2 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
													time.Sleep(1 * time.Second)
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}else {
												time.Sleep(1 * time.Second)
												hp2 = hp2 - attack2
												fmt.Println("Enemy took",attack2,"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
												}
											} else {
												time.Sleep(1 * time.Second)
												fmt.Println("attack missed")
												time.Sleep(1 * time.Second)
											}
										} else if a == 3 {
											if hp1 == hp1or {
												time.Sleep(1 * time.Second)
												fmt.Println("You're not low enough to heal")
												time.Sleep(1 * time.Second)
											} else if hp1 >= (hp1or - attack3){
												time.Sleep(1 * time.Second)
												hp1 = hp1or
												fmt.Println("You have been healed to max")
												time.Sleep(1 * time.Second)
											} else {
												time.Sleep(1 * time.Second)
												hp1 = hp1 + attack3
												fmt.Println("You have been healed",attack3,"health")
												time.Sleep(1 * time.Second)
											}
										} else if a == 4 {
											time.Sleep(1 * time.Second)
											fmt.Println("Which potion would you like to use")
											fmt.Println("1.Poison:",poip)
											fmt.Println("2.Curse",cursep)
											fmt.Println("3.Life Steal",hpst)
											fmt.Scanln(&b)
											time.Sleep(1 * time.Second)
											if b == 1 {
												if poip == 0 {
													time.Sleep(1 * time.Second)
													fmt.Println("Sorry you don't have this item")
													time.Sleep(1 * time.Second)
													continue
												} else {
													poip--
													time.Sleep(1 * time.Second)
													rpoip = 5
												}
											} else if b == 2 {
												if cursep == 0 {
													time.Sleep(1 * time.Second)
													fmt.Println("Sorry you don't have this item")
													time.Sleep(1 * time.Second)
													continue
												} else {
													cursep--
													time.Sleep(1 * time.Second)
													rcursep = 10
												}
											} else if b == 3 {
												if hpst == 0 {
													time.Sleep(1 * time.Second)
													fmt.Println("Sorry you don't have this item")
													time.Sleep(1 * time.Second)
													continue
												} else {
													hpst--
													if hp1 == hp1or {
														hp2 = hp2 - 30
														time.Sleep(1 * time.Second)
														fmt.Println("You're not low enough to heal")
														time.Sleep(1 * time.Second)
														fmt.Println("Enemy took 30 damage")
													} else if hp1 >= (hp1or - 30){
														time.Sleep(1 * time.Second)
														hp1 = hp1or
														hp2 = hp2 - 30
														fmt.Println("You have been healed to max")
														time.Sleep(1 * time.Second)
														fmt.Println("Enemy took 30 damage")
														if hp2 <= 0 {
															fmt.Println("You Win")
															break
														}
													} else {
														time.Sleep(1 * time.Second)
														hp1 = hp1 + 30
														hp2 = hp2 - 30
														fmt.Println("You stole 30 hp")
														if hp2 <= 0 {
															fmt.Println("You Win")
															break
														}
													}
												}
											}
										} else {
											fmt.Println("Sorry that is an invalid command")
										}
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(5)
										if random == 3 {
											if hp1 == hp1or / 2 {
												hp1 = hp1 + pethe
												time.Sleep(1 * time.Second)
												fmt.Println("Your pet healed you",pethe,"hp")
											} else {
												hp2 = hp2 - petat
												time.Sleep(1 * time.Second)
												fmt.Println("Your pet did",petat,"damage")
											}
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}
										time.Sleep(1 * time.Second)
										fmt.Println("enemy is attacking")
										time.Sleep(1 * time.Second)
										if hp2 <= (hp2or / 3) && hp1 > attacken2 {
											time.Sleep(1 * time.Second)
											hp2 = hp2 + attacken3
											fmt.Println("enemy has been healed",attacken3,"health")
											time.Sleep(1 * time.Second)
											continue
										}
										if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp1 = hp1 - attacken1 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2 
													break
												}	
											}else {
												fmt.Println("you took",attacken1,"damage")
												hp1 = hp1-attacken1
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2 
													break
												}
											}
										} else {
											var b = time.Now().UnixNano()
											rand.Seed(b)
											var random3 = rand.Intn(5)
												if random3 == 1 || random3 == 2 || random3 == 3 {
													var s = time.Now().UnixNano()
													rand.Seed(s)
													var random = rand.Intn(8)
													if random == 0 {
														hp1 = hp1 - attacken2 * 3 / 2
														fmt.Println("Critical Hit")
														time.Sleep(1 * time.Second)
														fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
														time.Sleep(1 * time.Second)
														if hp1 <= 0 {
															fmt.Println("You Died")
															check = 2 
															break
														}	
													}else {
														fmt.Println("you took",attacken2,"damage")
														hp1 = hp1-attacken2
														time.Sleep(1 * time.Second)	
														if hp1 <= 0 {
															fmt.Println("You Died")
															check = 2 
															break
														}
													}
												} else {
													fmt.Println("attack missed")
														time.Sleep(1 * time.Second)	
														continue
												}
											}
											if rpoip != 0 {
												hp2 = hp2 - 10
												rpoip--
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy poisoned 10 damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}
											if rcursep != 0 {
												if rcursep == 1 {
													hp2 = hp2 - 150
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took 150 damage.")
													time.Sleep(1 * time.Second)
													rcursep = 0
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}
												rcursep--
												time.Sleep(1 * time.Second)
												fmt.Println(rcursep,"turns left")
												time.Sleep(1 * time.Second)
											}
										}
										if check == 2 {
											time.Sleep(2 * time.Second)
											fmt.Println("Respawning")
											time.Sleep(5 * time.Second)
											continue
										}
									hp1 = hp1or
									attack1tr = 0
									hp1tr = 0
									attack2tr = 0
									attack3tr = 0
									attack4tr = 0
									montr = 0
									break
								}
								time.Sleep(1 * time.Second)
								fmt.Println("After Beating Hades, He drops down.")
								time.Sleep(5 * time.Second)
								fmt.Println("HOW!, he screams as he dies, but something is wrong, the curse hasn't worn off.")
								time.Sleep(5 * time.Second)
								fmt.Println("Right as you realize, the Whole castle comes crashing down and you fall over into darkness.")
								time.Sleep(5 * time.Second)
								for space := 1 ; space <=100 ; space++{
									fmt.Println()
								}
								fmt.Println("MAP 5 : TOADSTOOL VILLAGE")
								for space := 1 ; space <=20 ; space++{
									fmt.Println()
								}
								time.Sleep(2 * time.Second)	
								for space := 1 ; space <=100 ; space++{
									fmt.Println()
								}
								for {
									fmt.Println("You wake up, lost, in a forest, with no memory to recall.")
									time.Sleep(4 * time.Second)
									fmt.Println("The leaves of the trees, blocked the sunlight from the two suns.")
									time.Sleep(4 * time.Second)
									fmt.Println("You notice a sword in your hand, in horrible condition.") 
									time.Sleep(4 * time.Second)
									fmt.Println("Confused, you get up, despite the physical pain, You march forward in hope to find life, but it is hopeless.")
									time.Sleep(4 * time.Second)
									fmt.Println("Suddenly, you are shocked with an impact to the back of your head. It's a toad.")
									time.Sleep(4 * time.Second)
									fmt.Println("HAHAHAHAHHA! Laughs a sinister voice. My minion almost KILLED You! I'm TOAD MAN!")
									time.Sleep(4 * time.Second)
									fmt.Println("Toad man attacks right.")
									fmt.Println("1.Block left")
									fmt.Println("2.Block right")
									fmt.Println("3.Duck")
									fmt.Scanln(&b)
									time.Sleep(2 * time.Second)
										check = 1
										if b == 1 {
											fmt.Println("Toad man killed you")
											time.Sleep(2 * time.Second)
											fmt.Println("Respawing")
											time.Sleep(4 * time.Second)
											continue
										} else if b == 2 {
											fmt.Println("You blocked Toad man's attack")
										} else {
											fmt.Println("Toad man killed you")
											time.Sleep(2 * time.Second)
											fmt.Println("Respawing")
											time.Sleep(4 * time.Second)
											continue
										}
										time.Sleep(2 * time.Second)	
										fmt.Println("YOU WONT BEAT ME!")
										time.Sleep(2 * time.Second)	
										fmt.Println("Toad man uses lightning bolt.")
										fmt.Println("1.Block left")
										fmt.Println("2.Block right")
										fmt.Println("3.Roll")
										fmt.Scanln(&b)
										time.Sleep(2 * time.Second)	
										if b == 1 {
											fmt.Println("Toad man killed you")
											time.Sleep(2 * time.Second)
											fmt.Println("Respawing")
											time.Sleep(4 * time.Second)
											continue
										} else if b == 2 {
											fmt.Println("Toad man killed you")
											time.Sleep(2 * time.Second)
											fmt.Println("Respawing")
											time.Sleep(4 * time.Second)
											continue
										} else {
											fmt.Println("You rolled under the lightning bolt")
										}
										time.Sleep(2 * time.Second)	
										fmt.Println("EVIL SHALL RULE ALL!")
										time.Sleep(2 * time.Second)	
										fmt.Println("Toad Man Slashes Through middle!")
										fmt.Println("1.Parry")
										fmt.Println("2.Duck")
										fmt.Println("3.Block left")
										fmt.Scanln(&b)
										time.Sleep(2 * time.Second)	
										if b == 1 {
											fmt.Println("You dodged his attack")
										} else if b == 2 {
											fmt.Println("Toad man killed you")
											time.Sleep(2 * time.Second)
											fmt.Println("Respawing")
											time.Sleep(4 * time.Second)
											continue
										} else {
											fmt.Println("Toad man killed you")
											time.Sleep(2 * time.Second)
											fmt.Println("Respawing")
											time.Sleep(4 * time.Second)
											continue				
										}
										time.Sleep(2 * time.Second)	
										fmt.Println("Toad man uses and ariel attack. DIE!")
										fmt.Println("1.Uppercut")
										fmt.Println("2.Dive")
										fmt.Println("3.Parry")
										fmt.Scanln(&b)
										time.Sleep(2 * time.Second)	
										if b == 1 {
											fmt.Println("You punch Toad man in the face.")
										} else if b == 2 {
											fmt.Println("Toad man killed you")
											time.Sleep(2 * time.Second)
											fmt.Println("Respawing")
											time.Sleep(4 * time.Second)
											continue
										} else {
											fmt.Println("Toad man killed you")
											time.Sleep(2 * time.Second)
											fmt.Println("Respawing")
											time.Sleep(4 * time.Second)
											continue
										}
										break
								}
								time.Sleep(2 * time.Second)
								fmt.Println("he is quite hurt from all your counters, and he takes one last stand.")
								for {
									hp1 = hp1 - hp1tr 
									hp1 = hp1or
									attack1 = attack1 - attack1tr
									mon = mon - montr
									hp1 = hp1 - hp1tr
									hp1or = hp1
									attack2 = attack2 - attack2tr
									attack3 = attack3 - attack3tr
									attack4 = attack4 - attack4tr
									hp2or = hp2
									attack1tr = 0
									hp1tr = 0
									attack2tr = 0
									attack3tr = 0
									attack4tr = 0
									montr = 0
									check = 1
									rpoip = 0
									rcursep = 0
									hp2 = 800
									hp2or = hp2
									attacken1 = 80
									attacken2 = 120
									attacken3 = 50
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
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 3 {
													hp2 = hp2 - attack4 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}else {
													hp2 = hp2 - attack4
													fmt.Println("Enemy took",attack4,"damage")
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}
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
										fmt.Println()
										fmt.Println("choose your attack")
										fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
										fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
										fmt.Println("3.heal",attack3,"health 100% accuracy")
										fmt.Println("4.dark magic 100% accuracy")
										fmt.Scanln(&a)
										if a == 1 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp2 = hp2 - attack1 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}else {
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",attack1,"damage")
												hp2 = hp2 - attack1
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
											}
										} else if a == 2 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(3)
											if random == 1 || random == 2 || random == 3 {
												var s = time.Now().UnixNano()
												rand.Seed(s)
												var random = rand.Intn(8)
												if random == 3 {
													hp2 = hp2 - attack2 * 3 / 2
													fmt.Println("Critical Hit")
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
													time.Sleep(1 * time.Second)
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}else {
												time.Sleep(1 * time.Second)
												hp2 = hp2 - attack2
												fmt.Println("Enemy took",attack2,"damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
												}
											} else {
												time.Sleep(1 * time.Second)
												fmt.Println("attack missed")
												time.Sleep(1 * time.Second)
											}
										} else if a == 3 {
											if hp1 == hp1or {
												time.Sleep(1 * time.Second)
												fmt.Println("You're not low enough to heal")
												time.Sleep(1 * time.Second)
											} else if hp1 >= (hp1or - attack3){
												time.Sleep(1 * time.Second)
												hp1 = hp1or
												fmt.Println("You have been healed to max")
												time.Sleep(1 * time.Second)
											} else {
												time.Sleep(1 * time.Second)
												hp1 = hp1 + attack3
												fmt.Println("You have been healed",attack3,"health")
												time.Sleep(1 * time.Second)
											}
										} else if a == 4 {
											time.Sleep(1 * time.Second)
											fmt.Println("Which potion would you like to use")
											fmt.Println("1.Poison:",poip)
											fmt.Println("2.Curse",cursep)
											fmt.Println("3.Life Steal",hpst)
											fmt.Scanln(&b)
											time.Sleep(1 * time.Second)
											if b == 1 {
												if poip == 0 {
													time.Sleep(1 * time.Second)
													fmt.Println("Sorry you don't have this item")
													time.Sleep(1 * time.Second)
													continue
												} else {
													poip--
													time.Sleep(1 * time.Second)
													rpoip = 5
												}
											} else if b == 2 {
												if cursep == 0 {
													time.Sleep(1 * time.Second)
													fmt.Println("Sorry you don't have this item")
													time.Sleep(1 * time.Second)
													continue
												} else {
													cursep--
													time.Sleep(1 * time.Second)
													rcursep = 10
												}
											} else if b == 3 {
												if hpst == 0 {
													time.Sleep(1 * time.Second)
													fmt.Println("Sorry you don't have this item")
													time.Sleep(1 * time.Second)
													continue
												} else {
													hpst--
													if hp1 == hp1or {
														hp2 = hp2 - 30
														time.Sleep(1 * time.Second)
														fmt.Println("You're not low enough to heal")
														time.Sleep(1 * time.Second)
														fmt.Println("Enemy took 30 damage")
													} else if hp1 >= (hp1or - 30){
														time.Sleep(1 * time.Second)
														hp1 = hp1or
														hp2 = hp2 - 30
														fmt.Println("You have been healed to max")
														time.Sleep(1 * time.Second)
														fmt.Println("Enemy took 30 damage")
														if hp2 <= 0 {
															fmt.Println("You Win")
															break
														}
													} else {
														time.Sleep(1 * time.Second)
														hp1 = hp1 + 30
														hp2 = hp2 - 30
														fmt.Println("You stole 30 hp")
														if hp2 <= 0 {
															fmt.Println("You Win")
															break
														}
													}
												}
											}
										} else {
											fmt.Println("Sorry that is an invalid command")
										}
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(5)
										if random == 3 {
											if hp1 == hp1or / 2 {
												hp1 = hp1 + pethe
												time.Sleep(1 * time.Second)
												fmt.Println("You have been healed",pethe,"hp")
											} else {
												hp2 = hp2 - petat
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took",petat,"damage")
											}
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}
										time.Sleep(1 * time.Second)
										fmt.Println("enemy is attacking")
										time.Sleep(1 * time.Second)
										if hp2 <= (hp2or / 3) && hp1 > attacken2 {
											time.Sleep(1 * time.Second)
											hp2 = hp2 + attacken3
											fmt.Println("enemy has been healed",attacken3,"health")
											time.Sleep(1 * time.Second)
											continue
										}
										if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp1 = hp1 - attacken1 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2
													break
												}	
											}else {
												fmt.Println("you took",attacken1,"damage")
												hp1 = hp1-attacken1
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													check = 2
													break
												}
											}
										} else {
											var b = time.Now().UnixNano()
											rand.Seed(b)
											var random3 = rand.Intn(5)
												if random3 == 1 || random3 == 2 || random3 == 3 {
													var s = time.Now().UnixNano()
													rand.Seed(s)
													var random = rand.Intn(8)
													if random == 0 {
														hp1 = hp1 - attacken2 * 3 / 2
														fmt.Println("Critical Hit")
														time.Sleep(1 * time.Second)
														fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
														time.Sleep(1 * time.Second)
														if hp1 <= 0 {
															fmt.Println("You Died")
															check = 2
															break
														}	
													}else {
														fmt.Println("you took",attacken2,"damage")
														hp1 = hp1-attacken2
														time.Sleep(1 * time.Second)	
														if hp1 <= 0 {
															fmt.Println("You Died")
															check = 2
															break
														}
													}
												} else {
													fmt.Println("attack missed")
														time.Sleep(1 * time.Second)	
														continue
												}
											}
											if rpoip != 0 {
												hp2 = hp2 - 10
												rpoip--
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy poisoned 10 damage")
												time.Sleep(1 * time.Second)
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}	
											}
											if rcursep != 0 {
												if rcursep == 1 {
													hp2 = hp2 - 150
													time.Sleep(1 * time.Second)
													fmt.Println("Enemy took 150 damage.")
													time.Sleep(1 * time.Second)
													rcursep = 0
													if hp2 <= 0 {
														fmt.Println("You Win")
														break
													}	
												}
												rcursep--
												time.Sleep(1 * time.Second)
												fmt.Println(rcursep,"turns left")
												time.Sleep(1 * time.Second)
											}
										}
										if check == 2 {
											time.Sleep(2 * time.Second)
											fmt.Println("Respawning")
											time.Sleep(5 * time.Second)
											continue
										}
										hp1 = hp1or
										attack1tr = 0
										hp1tr = 0
										attack2tr = 0
										attack3tr = 0
										attack4tr = 0
										montr = 0
										break
								}
								fmt.Println("Toad man is weakened and falls to the ground. You place your sword at him but you see him grinning.")
								time.Sleep(5 * time.Second)
								fmt.Println("He jumps at you and knocks you over, breaking all your armor.")
								time.Sleep(5 * time.Second)
								fmt.Println("As a last resort you pull out the potion given to you by the shady merchant. You remember his words.")
								time.Sleep(5 * time.Second)
								fmt.Println("Use this to beat the one with the shroom.")
								time.Sleep(5 * time.Second)
								fmt.Println("Toad man marches closer and closer to you and you throw the potion at his head and he disintegrates into ashes.")
								time.Sleep(5 * time.Second)
								fmt.Println("The curse has been lifted, and you hike up the rocky trail. You place your blade next to you as you sit on a rock.")
								time.Sleep(5 * time.Second)
								fmt.Println("Your wolf lays next to you as the two suns set over Mystique Woods.")
								time.Sleep(7 * time.Second)
								for space := 1 ; space <=100 ; space++{
									fmt.Println()
								}
								fmt.Println("THE END")
								for space := 1 ; space <=20 ; space++{
									fmt.Println()
								}
								time.Sleep(4 * time.Second)	
								for space := 1 ; space <=100 ; space++{
									fmt.Println()
								}
								for space := 1 ; space <=100 ; space++{
									fmt.Println()
								}
								fmt.Println("Gerald: Coding")
								for space := 1 ; space <=20 ; space++{
									fmt.Println()
								}
								time.Sleep(4 * time.Second)	
								for space := 1 ; space <=100 ; space++{
									fmt.Println()
								}
								for space := 1 ; space <=100 ; space++{
									fmt.Println()
								}
								fmt.Println("Aarez: Visual and story development")
								for space := 1 ; space <=20 ; space++{
									fmt.Println()
								}
								time.Sleep(4 * time.Second)	
								for space := 1 ; space <=100 ; space++{
									fmt.Println()
								}
								for space := 1 ; space <=100 ; space++{
									fmt.Println()
								}
								fmt.Println("Jonathan: Enemy Info, and Testing")
								for space := 1 ; space <=20 ; space++{
									fmt.Println()
								}
								time.Sleep(4 * time.Second)	
								for space := 1 ; space <=100 ; space++{
									fmt.Println()
								}
								for space := 1 ; space <=100 ; space++{
									fmt.Println()
								}
								fmt.Println("Specail thanks to: Hagan, Suleman, Alex, Zaid, Huzaifa, and Bryce")
								for space := 1 ; space <=20 ; space++{
									fmt.Println()
								}
								time.Sleep(4 * time.Second)	
								for space := 1 ; space <=100 ; space++{
									fmt.Println()
								}
								for space := 1 ; space <=100 ; space++{
									fmt.Println()
								}
								fmt.Println("DEMON'S GATE")
								for space := 1 ; space <=20 ; space++{
									fmt.Println()
								}
								time.Sleep(4 * time.Second)	
								for space := 1 ; space <=100 ; space++{
									fmt.Println()
								}
			}else if game == 3 {
				for {
					var e int
					var c int
					var d int
					var p1hp = 100
					var p1hpor = p1hp
					var p1attack1 = 10
					var p1attack2 = 30
					var p1attack3 = 20
					var p2hp = 100
					var p2hpor = p2hp
					var p2attack1 = 10
					var p2attack2 = 30
					var p2attack3 = 20
					var p1mon = 300
					var p2mon = 300
					fmt.Println("Player 1 is buying supplies")
					time.Sleep(2 * time.Second)
					for {
						fmt.Println("What would you like to buy")
						fmt.Println("Total Coins:",p1mon)
						fmt.Println("1.Armor Upgrade: 50 coins")
						fmt.Println("2.Weapon Upgrade: 50 coins")
						fmt.Println("3.Magic Upgrade: 50 coins")
						fmt.Println("4.Exit")
						fmt.Scanln(&e)
						if e == 1 {
							if p1mon < 50{
								time.Sleep(1 * time.Second)
								fmt.Println("You do not have enough to buy this item")
								time.Sleep(1 * time.Second)
							}else{
								p1mon = p1mon - 50
								p1hp = p1hp + 25
								p1hpor = p1hp
								time.Sleep(1 * time.Second)
								fmt.Println("Armor has been upgraded")
								time.Sleep(1 * time.Second)
							}
						}else if e == 2 {
							if p1mon < 50{
								time.Sleep(1 * time.Second)
								fmt.Println("You do not have enough to buy this item")
								time.Sleep(1 * time.Second)
							}else{
								p1mon = p1mon - 50
								p1attack1 = p1attack1 + 10
								p1attack2 = p1attack2 + 10
								time.Sleep(1 * time.Second)
								fmt.Println("Weapons has been upgraded")
								time.Sleep(1 * time.Second)
							}
						}else if e == 3{
							if p1mon < 50{
								time.Sleep(1 * time.Second)
								fmt.Println("You do not have enough to buy this item")
								time.Sleep(1 * time.Second)
							}else{
								p1mon = p1mon - 50
								p1attack3 = p1attack3 + 10
								time.Sleep(1 * time.Second)
								fmt.Println("Magic has been upgraded")
								time.Sleep(1 * time.Second)
							}
						}else{
							time.Sleep(1 * time.Second)
							break
						}
					}
					time.Sleep(2 * time.Second)
					fmt.Println("Player 2 is buying suplies")
					time.Sleep(2 * time.Second)
					for {
						fmt.Println("What would you like to buy")
						fmt.Println("Total Coins:",p2mon)
						fmt.Println("1.Armor Upgrade: 50 coins")
						fmt.Println("2.Weapon Upgrade: 50 coins")
						fmt.Println("3.Magic Upgrade: 50 coins")
						fmt.Println("4.Exit")
						fmt.Scanln(&e)
						if e == 1 {
							if p2mon < 50{
								time.Sleep(1 * time.Second)
								fmt.Println("You do not have enough to buy this item")
								time.Sleep(1 * time.Second)
							}else{
								p2mon = p2mon - 50
								p2hp = p2hp + 25
								p2hpor = p2hp
								time.Sleep(1 * time.Second)
								fmt.Println("Armor has been upgraded")
								time.Sleep(1 * time.Second)
							}
						}else if e == 2 {
							if p2mon < 50{
								time.Sleep(1 * time.Second)
								fmt.Println("You do not have enough to buy this item")
								time.Sleep(1 * time.Second)
							}else{
								p2mon = p2mon - 50
								p2attack1 = p2attack1 + 10
								p2attack2 = p2attack2 + 10
								time.Sleep(1 * time.Second)
								fmt.Println("Weapons has been upgraded")
								time.Sleep(1 * time.Second)
							}
						}else if e == 3{
							if p2mon < 50{
								time.Sleep(1 * time.Second)
								fmt.Println("You do not have enough to buy this item")
								time.Sleep(1 * time.Second)
							}else{
								p2mon = p2mon - 50
								p2attack3 = p2attack3 + 10
								time.Sleep(1 * time.Second)
								fmt.Println("Magic has been upgraded")
								time.Sleep(1 * time.Second)
							}
						}else{
							time.Sleep(1 * time.Second)
							break
						}
					}
					fmt.Println("Battle is starting")
					time.Sleep(2 * time.Second)
					fmt.Println("Player 1 hp",p1hp,"          Player 2 hp",p2hp)
					for battle1 := 1 ;; battle1++ {
						time.Sleep(2 * time.Second)
						fmt.Println("Player 1 is attacking")
						time.Sleep(2 * time.Second)
						fmt.Println("Player 1 hp",p1hp,"          Player 2 hp",p2hp)
						fmt.Println()
						fmt.Println("choose your attack")
						fmt.Println("1.light attack ",p1attack1,"damage 100% accuracy")
						fmt.Println("2.heavy attack ",p1attack2,"damage 75% accuracy")
						fmt.Println("3.heal",p1attack3,"health 100% accuracy")
						fmt.Scanln(&d)
						if d == 1 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(8)
							if random == 0 {
								p2hp = p2hp - p2attack1 * 3 / 2
								fmt.Println("Critical Hit")
								time.Sleep(1 * time.Second)
								fmt.Println("Player 2 took",p1attack1 + (p1attack1 / 2),"damage")
								time.Sleep(1 * time.Second)
								if p2hp <= 0 {
									fmt.Println("Player 1 Wins")
									break
								}
							}else {
								time.Sleep(1 * time.Second)
								fmt.Println("Player 2 took",p1attack1,"damage")
								p2hp = p2hp - p1attack1
								time.Sleep(1 * time.Second)
								if p2hp <= 0 {
									fmt.Println("Player 1 Wins")
									break
								}
							}
						} else if d == 2 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(3)
							if random == 1 || random == 2 || random == 3 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 3 {
									p2hp = p2hp - p1attack2 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("Player 2 took",p1attack2 + (p1attack2 / 2),"damage")
									time.Sleep(1 * time.Second)
									if p2hp <= 0 {
										fmt.Println("Player 1 Wins")
										break
									}	
								} else {
									time.Sleep(1 * time.Second)
									p2hp = p2hp - p1attack2
								fmt.Println("Player 2 took",p1attack2,"damage")
								time.Sleep(1 * time.Second)
								if p2hp <= 0 {
									fmt.Println("Player 1 Wins")
									break
								}
								}
							} else {
								time.Sleep(1 * time.Second)
								fmt.Println("attack missed")
								time.Sleep(1 * time.Second)
							}
						} else if d == 3 {
							if p1hp == p1hpor {
								time.Sleep(1 * time.Second)
								fmt.Println("You're not low enough to heal")
								time.Sleep(1 * time.Second)
							} else if p1hp >= (p1hpor - p1attack3){
								time.Sleep(1 * time.Second)
								p1hp = p1hpor
								fmt.Println("You have been healed to max")
								time.Sleep(1 * time.Second)
							} else {
								time.Sleep(1 * time.Second)
								p1hp = p1hp + p1attack3
								fmt.Println("You have been healed",p1attack3,"health")
								time.Sleep(1 * time.Second)
							}
						} else {
							fmt.Println("Sorry that is an invalid command")
						}
						time.Sleep(2 * time.Second)
						fmt.Println("Player 2 is attacking")
						time.Sleep(2 * time.Second)
						fmt.Println("Player 1 hp",p1hp,"          Player 2 hp",p2hp)
						fmt.Println()
						fmt.Println("choose your attack")
						fmt.Println("1.light attack ",p2attack1,"damage 100% accuracy")
						fmt.Println("2.heavy attack ",p2attack2,"damage 75% accuracy")
						fmt.Println("3.heal",p2attack3,"health 100% accuracy")
						fmt.Scanln(&d)
						if d == 1 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(8)
							if random == 0 {
								p1hp = p1hp - p2attack1 * 3 / 2
								fmt.Println("Critical Hit")
								time.Sleep(1 * time.Second)
								fmt.Println("Player 1 took",p2attack1 + (p2attack1 / 2),"damage")
								time.Sleep(1 * time.Second)
								if p1hp <= 0 {
									fmt.Println("Player 2 Wins")
									break
								}
							}else {
								time.Sleep(1 * time.Second)
								fmt.Println("Player 1 took",p2attack1,"damage")
								p1hp = p1hp - p2attack1
								time.Sleep(1 * time.Second)
								if p1hp <= 0 {
									fmt.Println("Player 2 Wins")
									break
								}
							}
						} else if d == 2 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(3)
							if random == 1 || random == 2 || random == 3 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 3 {
									p1hp = p1hp - p2attack2 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("Player 2 took",p2attack2 + (p2attack2 / 2),"damage")
									time.Sleep(1 * time.Second)
									if p1hp <= 0 {
										fmt.Println("Player 2 Wins")
										break
									}	
								} else {
										time.Sleep(1 * time.Second)
										p1hp = p1hp - p2attack2
									fmt.Println("Player 1 took",p2attack2,"damage")
									time.Sleep(1 * time.Second)
									if p1hp <= 0 {
										fmt.Println("Player 2 Wins")
										break
									}
								}
							} else {
								time.Sleep(1 * time.Second)
								fmt.Println("attack missed")
								time.Sleep(1 * time.Second)
							}
						} else if d == 3 {
							if p2hp == p2hpor {
								time.Sleep(1 * time.Second)
								fmt.Println("You're not low enough to heal")
								time.Sleep(1 * time.Second)
							} else if p2hp >= (p2hpor - p2attack3){
								time.Sleep(1 * time.Second)
								p2hp = p2hpor
								fmt.Println("You have been healed to max")
								time.Sleep(1 * time.Second)
							} else {
								time.Sleep(1 * time.Second)
								p2hp = p2hp + p2attack3
								fmt.Println("You have been healed",p2attack3,"health")
								time.Sleep(1 * time.Second)
							}
						} else {
							fmt.Println("Sorry that is an invalid command")
						}
						}
						time.Sleep(3 * time.Second)
						fmt.Println("Would you like to play again:")
						fmt.Println("1.Yes")
						fmt.Println("2.No")
						fmt.Scanln(&c)
						if c == 1 {
							time.Sleep(3 * time.Second)
							fmt.Println("Starting new game")
							time.Sleep(3 * time.Second)
						} else {
							time.Sleep(3 * time.Second)
							break
						}
				}
			}else if game == 4 {
				for {
					var check int
					var a int
					var attack1 = 10
					var attack2 = 30
					var attack3 = 20
					var attack4 = 50
					var b int
					var mon = 500
					var hp1 = 100
					var	hp1or = hp1
					var hp2 = 70
					var	hp2or = hp2
					var attacken1 = 5
					var attacken2 = 15
					var attacken3 = 5
					fmt.Println("Welcome to Survival")
					time.Sleep(3 * time.Second)
					fmt.Println("Buying supplies")
					time.Sleep(3 * time.Second)
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
					fmt.Println("Round:1")
					for sur := 2 ;; sur++{
						attacken1 = attacken1 + 5
						attacken2 = attacken2 + 5
						attacken3 = attacken3 + 5
						hp2 = hp2 + 5
						hp2or = hp2
						time.Sleep(2 * time.Second)	
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
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 3 {
										hp2 = hp2 - attack4 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									}else {
										hp2 = hp2 - attack4
										fmt.Println("Enemy took",attack4,"damage")
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
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
							fmt.Println()
							fmt.Println("choose your attack")
							fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
							fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
							fmt.Println("3.heal",attack3,"health 100% accuracy")
							fmt.Scanln(&a)
							if a == 1 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 0 {
									hp2 = hp2 - attack1 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}else {
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack1,"damage")
									hp2 = hp2-attack1
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}
								}
							} else if a == 2 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(3)
								if random == 1 || random == 2 || random == 3 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 3 {
										hp2 = hp2 - attack2 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									}else {
									time.Sleep(1 * time.Second)
									hp2 = hp2 - attack2
									fmt.Println("Enemy took",attack2,"damage")
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}
									}
								} else {
									time.Sleep(1 * time.Second)
									fmt.Println("attack missed")
									time.Sleep(1 * time.Second)
								}
							} else if a == 3 {
								if hp1 == hp1or {
									time.Sleep(1 * time.Second)
									fmt.Println("You're not low enough to heal")
									time.Sleep(1 * time.Second)
								} else if hp1 >= (hp1or - attack3){
									time.Sleep(1 * time.Second)
									hp1 = hp1or
									fmt.Println("You have been healed to max")
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
							time.Sleep(1 * time.Second)
							fmt.Println("enemy is attacking")
							time.Sleep(1 * time.Second)
							if hp2 <= (hp2or / 2) && hp1 > attacken2 {
								time.Sleep(1 * time.Second)
								hp2 = hp2 + attacken3
								fmt.Println("enemy has been healed",attacken3,"health")
								time.Sleep(1 * time.Second)
								continue
							}
							if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 0 {
									hp1 = hp1 - attacken1 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
									time.Sleep(1 * time.Second)
									if hp1 <= 0 {
										fmt.Println("You Died")
										time.Sleep(1 * time.Second)	
										fmt.Println("Congragulations! You lasted",sur - 1,"rounds")
										time.Sleep(4 * time.Second)	
										check = 2
										break
									}	
								} else {
									fmt.Println("you took",attacken1,"damage")
									hp1 = hp1-attacken1
									time.Sleep(1 * time.Second)
									if hp1 <= 0 {
										fmt.Println("You Died")
										time.Sleep(1 * time.Second)	
										fmt.Println("Congragulations! You lasted",sur - 1,"rounds")
										time.Sleep(4 * time.Second)	
										check = 2
										break
									}
								}
							} else {
								var b = time.Now().UnixNano()
								rand.Seed(b)
								var random3 = rand.Intn(5)
									if random3 == 1 || random3 == 2 || random3 == 3 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 0 {
											hp1 = hp1 - attacken2 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp1 <= 0 {
												fmt.Println("You Died")
												time.Sleep(1 * time.Second)	
												fmt.Println("Congragulations! You lasted",sur - 1,"rounds")
												time.Sleep(4 * time.Second)	
												check = 2
												break
											}	
										}else {
											fmt.Println("you took",attacken2,"damage")
											hp1 = hp1-attacken2
											time.Sleep(1 * time.Second)	
											if hp1 <= 0 {
												fmt.Println("You Died")
												check = 2
												time.Sleep(1 * time.Second)	
												fmt.Println("Congragulations! You lasted",sur - 1,"rounds")
												time.Sleep(4 * time.Second)	
												break
											}
										}
									} else {
										fmt.Println("attack missed")
											time.Sleep(1 * time.Second)	
											continue
									}
								}
							}
							hp2 = hp2or
							time.Sleep(1 * time.Second)	
							if check == 2 {
								time.Sleep(1 * time.Second)	
								break
							} else {
								time.Sleep(1 * time.Second)	
								fmt.Println("Round:",sur)
								continue
							}
					}
					fmt.Println("Would you like to play again")
					fmt.Println("1.Yes")
					fmt.Println("2.No")
					fmt.Scanln(&a)
					if a == 1 {
						fmt.Println("Starting new round")
						time.Sleep(3 * time.Second)	
						continue
					} else {
						time.Sleep(3 * time.Second)	
						break
					}
				}
			} else if game == 5{
					var petat = 20
					var pethe = 20
					var poip = 1
					var cursep = 1
					var hpst = 1
					var rpoip = 0
					var rcursep = 0
					var a int
				  	//var check = 1
					var b int
					//var key = 1
					//var key2 = 1
					//var key3 = 1
					var mon int
					var hp1 = 10000
					var hp2 = 69
					var hp2or = hp2
					var hp1or = hp1
					var attack1 = 1000
					var attack2 = 30
					var attack3 = 20
					var attack4 = 50
					//var hp1tr int
					//var attack1tr int
					//var attack2tr int
					//var attack3tr int
					//var attack4tr int
					var attacken1 = 10
					var attacken2 = 20 
					var attacken3 = 10
					//var montr int
					//var check4 int
				fmt.Println("After defeating the lord of hell Toad Man, you were told that your memories would return.")
				time.Sleep(5 * time.Second)
				fmt.Println("The problem is you still do not have any memories before Mystic Woods.")
				time.Sleep(5 * time.Second)
				fmt.Println("You continue on your adventure despite this, but the first choice is the one needed for now and right now you have 2 places to go.")
				fmt.Println("1.Return to Mystic Woods")
				fmt.Println("2.Go up Mount Olympus")
				fmt.Scanln(&a)
				time.Sleep(2 * time.Second)
				fmt.Println("You have made your choice and you start a new quest, a quest reborn.")
				time.Sleep(8 * time.Second)
				for space := 1 ; space <=100 ; space++{
					fmt.Println()
				}
				fmt.Print("DEMON'S GATE")
				fmt.Println(" 2")
				fmt.Println("A Hero Reborn")
				for space := 1 ; space <=20 ; space++{
					fmt.Println()
				}
				time.Sleep(4 * time.Second)	
				for space := 1 ; space <=100 ; space++{
					fmt.Println()
				}
				time.Sleep(4 * time.Second)
				for space := 1 ; space <=100 ; space++{
					fmt.Println()
				}
				fmt.Print("Act I: A Hero Returns")
				for space := 1 ; space <=20 ; space++{
					fmt.Println()
				}
				time.Sleep(4 * time.Second)	
				for space := 1 ; space <=100 ; space++{
					fmt.Println()
				}
				time.Sleep(4 * time.Second)
				if a == 1 {
					fmt.Println("You return to the place that began your journey, the place you fought monsters, and the time you freed the woods.")
					time.Sleep(4 * time.Second)
					fmt.Println("It's been 2 years since that things have changed now, the underworld now controls the land and you have to free it again.")
					time.Sleep(4 * time.Second)
					fmt.Println("You continue forwards but find yourself upon enemy territory.")
					time.Sleep(4 * time.Second)
					fmt.Println("A cursed wood dweller stands upon you, with your rusty sword held in front of you, and starts the battle.")
					time.Sleep(4 * time.Second)
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
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 3 {
									hp2 = hp2 - attack4 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}else {
									hp2 = hp2 - attack4
									fmt.Println("Enemy took",attack4,"damage")
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}
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
						fmt.Println()
						fmt.Println("choose your attack")
						fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
						fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
						fmt.Println("3.heal",attack3,"health 100% accuracy")
						fmt.Println("4.dark magic 100% accuracy")
						fmt.Scanln(&a)
						if a == 1 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(8)
							if random == 0 {
								hp2 = hp2 - attack1 * 3 / 2
								fmt.Println("Critical Hit")
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}	
							}else {
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy took",attack1,"damage")
								hp2 = hp2 - attack1
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}
							}
						} else if a == 2 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(3)
							if random == 1 || random == 2 || random == 3 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 3 {
									hp2 = hp2 - attack2 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}else {
								time.Sleep(1 * time.Second)
								hp2 = hp2 - attack2
								fmt.Println("Enemy took",attack2,"damage")
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}
								}
							} else {
								time.Sleep(1 * time.Second)
								fmt.Println("attack missed")
								time.Sleep(1 * time.Second)
							}
						} else if a == 3 {
							if hp1 == hp1or {
								time.Sleep(1 * time.Second)
								fmt.Println("You're not low enough to heal")
								time.Sleep(1 * time.Second)
							} else if hp1 >= (hp1or - attack3){
								time.Sleep(1 * time.Second)
								hp1 = hp1or
								fmt.Println("You have been healed to max")
								time.Sleep(1 * time.Second)
							} else {
								time.Sleep(1 * time.Second)
								hp1 = hp1 + attack3
								fmt.Println("You have been healed",attack3,"health")
								time.Sleep(1 * time.Second)
							}
						} else if a == 4 {
							time.Sleep(1 * time.Second)
							fmt.Println("Which potion would you like to use")
							fmt.Println("1.Poison:",poip)
							fmt.Println("2.Curse",cursep)
							fmt.Println("3.Life Steal",hpst)
							fmt.Scanln(&b)
							time.Sleep(1 * time.Second)
							if b == 1 {
								if poip == 0 {
									time.Sleep(1 * time.Second)
									fmt.Println("Sorry you don't have this item")
									time.Sleep(1 * time.Second)
									continue
								} else {
									poip--
									time.Sleep(1 * time.Second)
									rpoip = 5
								}
							} else if b == 2 {
								if cursep == 0 {
									time.Sleep(1 * time.Second)
									fmt.Println("Sorry you don't have this item")
									time.Sleep(1 * time.Second)
									continue
								} else {
									cursep--
									time.Sleep(1 * time.Second)
									rcursep = 10
								}
							} else if b == 3 {
								if hpst == 0 {
									time.Sleep(1 * time.Second)
									fmt.Println("Sorry you don't have this item")
									time.Sleep(1 * time.Second)
									continue
								} else {
									hpst--
									if hp1 == hp1or {
										hp2 = hp2 - 30
										time.Sleep(1 * time.Second)
										fmt.Println("You're not low enough to heal")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took 30 damage")
									} else if hp1 >= (hp1or - 30){
										time.Sleep(1 * time.Second)
										hp1 = hp1or
										hp2 = hp2 - 30
										fmt.Println("You have been healed to max")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took 30 damage")
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
									} else {
										time.Sleep(1 * time.Second)
										hp1 = hp1 + 30
										hp2 = hp2 - 30
										fmt.Println("You stole 30 hp")
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
									}
								}
							}
						} else {
							fmt.Println("Sorry that is an invalid command")
						}
						var s = time.Now().UnixNano()
						rand.Seed(s)
						var random = rand.Intn(5)
						if random == 3 {
							if hp1 == hp1or / 2 {
								hp1 = hp1 + pethe
								time.Sleep(1 * time.Second)
								fmt.Println("You have been healed",pethe,"hp")
							} else {
								hp2 = hp2 - petat
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy took",petat,"damage")
							}
							if hp2 <= 0 {
								fmt.Println("You Win")
								break
							}	
						}
						time.Sleep(1 * time.Second)
						fmt.Println("enemy is attacking")
						time.Sleep(1 * time.Second)
						if hp2 <= (hp2or / 3) && hp1 > attacken2 {
							time.Sleep(1 * time.Second)
							hp2 = hp2 + attacken3
							fmt.Println("enemy has been healed",attacken3,"health")
							time.Sleep(1 * time.Second)
							continue
						}
						if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(8)
							if random == 0 {
								hp1 = hp1 - attacken1 * 3 / 2
								fmt.Println("Critical Hit")
								time.Sleep(1 * time.Second)
								fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
								time.Sleep(1 * time.Second)
								if hp1 <= 0 {
									fmt.Println("You Died")
									//check = 2
									break
								}	
							}else {
								fmt.Println("you took",attacken1,"damage")
								hp1 = hp1-attacken1
								time.Sleep(1 * time.Second)
								if hp1 <= 0 {
									fmt.Println("You Died")
									//check = 2
									break
								}
							}
						} else {
							var b = time.Now().UnixNano()
							rand.Seed(b)
							var random3 = rand.Intn(5)
								if random3 == 1 || random3 == 2 || random3 == 3 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp1 = hp1 - attacken2 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											//check = 2
											break
										}	
									}else {
										fmt.Println("you took",attacken2,"damage")
										hp1 = hp1-attacken2
										time.Sleep(1 * time.Second)	
										if hp1 <= 0 {
											fmt.Println("You Died")
											//check = 2
											break
										}
									}
								} else {
									fmt.Println("attack missed")
										time.Sleep(1 * time.Second)	
										continue
								}
							}
							if rpoip != 0 {
								hp2 = hp2 - 10
								rpoip--
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy poisoned 10 damage")
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}	
							}
							if rcursep != 0 {
								if rcursep == 1 {
									hp2 = hp2 - 150
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took 150 damage.")
									time.Sleep(1 * time.Second)
									rcursep = 0
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}
								rcursep--
								time.Sleep(1 * time.Second)
								fmt.Println(rcursep,"turns left")
								time.Sleep(1 * time.Second)
							}
						}
						hp1or = hp1or + 20
						hp1 = hp1or
						time.Sleep(1 * time.Second)
						fmt.Println("You gained 100 coins")
						time.Sleep(1 * time.Second)
						mon = mon + 100
						fmt.Println("Total coins:",mon)
						time.Sleep(4 * time.Second)
						fmt.Println("After easily defeating the Wood Dweller you grab his helmet and continue forwards.")
						time.Sleep(4 * time.Second)
						fmt.Println("Armor upraded")
						time.Sleep(4 * time.Second)
						fmt.Println("You know what your goal is, You will have to find the demon in control and defeat him.")
						time.Sleep(4 * time.Second)
						fmt.Println("While you trudge through the bushes you find two routes to choose from.")
						time.Sleep(4 * time.Second)
						fmt.Println("Either you can climb down a cliff or follow a creek.")
						time.Sleep(2 * time.Second) 
						fmt.Println("1.Creek")
						fmt.Println("2.Cliff")
						fmt.Scanln(&b)
						time.Sleep(2 * time.Second) 
						if b == 1 {
							fmt.Println("You follow the water down to a small pond, but a small Sea Monster ambushes you. Small Sea monster")
							time.Sleep(2 * time.Second) 
							rpoip = 0
							rcursep = 0
							hp2 = 100
							hp2or = hp2
							attacken1 = 10
							attacken2 = 30
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
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 3 {
											hp2 = hp2 - attack4 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
											hp2 = hp2 - attack4
											fmt.Println("Enemy took",attack4,"damage")
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
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
								fmt.Println()
								fmt.Println("choose your attack")
								fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
								fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
								fmt.Println("3.heal",attack3,"health 100% accuracy")
								fmt.Println("4.dark magic 100% accuracy")
								fmt.Scanln(&a)
								if a == 1 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp2 = hp2 - attack1 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									} else {
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack1,"damage")
										hp2 = hp2 - attack1
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
									}
								} else if a == 2 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(3)
									if random == 1 || random == 2 || random == 3 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 3 {
											hp2 = hp2 - attack2 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
										time.Sleep(1 * time.Second)
										hp2 = hp2 - attack2
										fmt.Println("Enemy took",attack2,"damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
										}
									} else {
										time.Sleep(1 * time.Second)
										fmt.Println("attack missed")
										time.Sleep(1 * time.Second)
									}
								} else if a == 3 {
									if hp1 == hp1or {
										time.Sleep(1 * time.Second)
										fmt.Println("You're not low enough to heal")
										time.Sleep(1 * time.Second)
									} else if hp1 >= (hp1or - attack3){
										time.Sleep(1 * time.Second)
										hp1 = hp1or
										fmt.Println("You have been healed to max")
										time.Sleep(1 * time.Second)
									} else {
										time.Sleep(1 * time.Second)
										hp1 = hp1 + attack3
										fmt.Println("You have been healed",attack3,"health")
										time.Sleep(1 * time.Second)
									}
								} else if a == 4 {
									time.Sleep(1 * time.Second)
									fmt.Println("Which potion would you like to use")
									fmt.Println("1.Poison:",poip)
									fmt.Println("2.Curse",cursep)
									fmt.Println("3.Life Steal",hpst)
									fmt.Scanln(&b)
									time.Sleep(1 * time.Second)
									if b == 1 {
										if poip == 0 {
											time.Sleep(1 * time.Second)
											fmt.Println("Sorry you don't have this item")
											time.Sleep(1 * time.Second)
											continue
										} else {
											poip--
											time.Sleep(1 * time.Second)
											rpoip = 5
										}
									} else if b == 2 {
										if cursep == 0 {
											time.Sleep(1 * time.Second)
											fmt.Println("Sorry you don't have this item")
											time.Sleep(1 * time.Second)
											continue
										} else {
											cursep--
											time.Sleep(1 * time.Second)
											rcursep = 10
										}
									} else if b == 3 {
										if hpst == 0 {
											time.Sleep(1 * time.Second)
											fmt.Println("Sorry you don't have this item")
											time.Sleep(1 * time.Second)
											continue
										} else {
											hpst--
											if hp1 == hp1or {
												hp2 = hp2 - 30
												time.Sleep(1 * time.Second)
												fmt.Println("You're not low enough to heal")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took 30 damage")
											} else if hp1 >= (hp1or - 30){
												time.Sleep(1 * time.Second)
												hp1 = hp1or
												hp2 = hp2 - 30
												fmt.Println("You have been healed to max")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took 30 damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
											} else {
												time.Sleep(1 * time.Second)
												hp1 = hp1 + 30
												hp2 = hp2 - 30
												fmt.Println("You stole 30 hp")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
											}
										}
									}
								} else {
									fmt.Println("Sorry that is an invalid command")
								}
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(5)
								if random == 3 {
									if hp1 == hp1or / 2 {
										hp1 = hp1 + pethe
										time.Sleep(1 * time.Second)
										fmt.Println("You have been healed",pethe,"hp")
									} else {
										hp2 = hp2 - petat
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",petat,"damage")
									}
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}
								time.Sleep(1 * time.Second)
								fmt.Println("enemy is attacking")
								time.Sleep(1 * time.Second)
								if hp2 <= (hp2or / 3) && hp1 > attacken2 {
									time.Sleep(1 * time.Second)
									hp2 = hp2 + attacken3
									fmt.Println("enemy has been healed",attacken3,"health")
									time.Sleep(1 * time.Second)
									continue
								}
								if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp1 = hp1 - attacken1 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											//check = 2
											break
										}	
									}else {
										fmt.Println("you took",attacken1,"damage")
										hp1 = hp1-attacken1
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											//check = 2
											break
										}
									}
								} else {
									var b = time.Now().UnixNano()
									rand.Seed(b)
									var random3 = rand.Intn(5)
										if random3 == 1 || random3 == 2 || random3 == 3 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp1 = hp1 - attacken2 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													//check = 2
													break
												}	
											}else {
												fmt.Println("you took",attacken2,"damage")
												hp1 = hp1-attacken2
												time.Sleep(1 * time.Second)	
												if hp1 <= 0 {
													fmt.Println("You Died")
													//check = 2
													break
												}
											}
										} else {
											fmt.Println("attack missed")
												time.Sleep(1 * time.Second)	
												continue
										}
									}
									if rpoip != 0 {
										hp2 = hp2 - 10
										rpoip--
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy poisoned 10 damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									}
									if rcursep != 0 {
										if rcursep == 1 {
											hp2 = hp2 - 150
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took 150 damage.")
											time.Sleep(1 * time.Second)
											rcursep = 0
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}
										rcursep--
										time.Sleep(1 * time.Second)
										fmt.Println(rcursep,"turns left")
										time.Sleep(1 * time.Second)
									}
								}
								hp1 = hp1or
								time.Sleep(1 * time.Second)
								fmt.Println("You gained 50 coins")
								time.Sleep(1 * time.Second)
								mon = mon + 50
								fmt.Println("Total coins:",mon)
								time.Sleep(4 * time.Second)
								fmt.Println("The monster falls and you gain a sea orb, you can either use it for a weapon upgrade or magic, which do you choose?")
								time.Sleep(2 * time.Second)
								fmt.Println("1.Weapon")
								fmt.Println("2.Magic")
								fmt.Scanln(&b)
								time.Sleep(2 * time.Second)
								if b == 1 {
									attack1 = attack1 + 20
									attack2 = attack2 + 20
									attack4 = attack4 + 20
									fmt.Println("Weapon upgraded")
								} else {
									attack3 = attack3 + 10
									fmt.Println("Magic upgraded")
								}
								time.Sleep(2 * time.Second)
						} else {
							fmt.Println("You scale the edge of the cliff and find yourself upon a cave spider")
							rpoip = 0
							rcursep = 0
							hp2 = 100
							hp2or = hp2
							attacken1 = 10
							attacken2 = 30
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
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 3 {
											hp2 = hp2 - attack4 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
											hp2 = hp2 - attack4
											fmt.Println("Enemy took",attack4,"damage")
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
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
								fmt.Println()
								fmt.Println("choose your attack")
								fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
								fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
								fmt.Println("3.heal",attack3,"health 100% accuracy")
								fmt.Println("4.dark magic 100% accuracy")
								fmt.Scanln(&a)
								if a == 1 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp2 = hp2 - attack1 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									} else {
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack1,"damage")
										hp2 = hp2 - attack1
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
									}
								} else if a == 2 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(3)
									if random == 1 || random == 2 || random == 3 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 3 {
											hp2 = hp2 - attack2 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
										time.Sleep(1 * time.Second)
										hp2 = hp2 - attack2
										fmt.Println("Enemy took",attack2,"damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
										}
									} else {
										time.Sleep(1 * time.Second)
										fmt.Println("attack missed")
										time.Sleep(1 * time.Second)
									}
								} else if a == 3 {
									if hp1 == hp1or {
										time.Sleep(1 * time.Second)
										fmt.Println("You're not low enough to heal")
										time.Sleep(1 * time.Second)
									} else if hp1 >= (hp1or - attack3){
										time.Sleep(1 * time.Second)
										hp1 = hp1or
										fmt.Println("You have been healed to max")
										time.Sleep(1 * time.Second)
									} else {
										time.Sleep(1 * time.Second)
										hp1 = hp1 + attack3
										fmt.Println("You have been healed",attack3,"health")
										time.Sleep(1 * time.Second)
									}
								} else if a == 4 {
									time.Sleep(1 * time.Second)
									fmt.Println("Which potion would you like to use")
									fmt.Println("1.Poison:",poip)
									fmt.Println("2.Curse",cursep)
									fmt.Println("3.Life Steal",hpst)
									fmt.Scanln(&b)
									time.Sleep(1 * time.Second)
									if b == 1 {
										if poip == 0 {
											time.Sleep(1 * time.Second)
											fmt.Println("Sorry you don't have this item")
											time.Sleep(1 * time.Second)
											continue
										} else {
											poip--
											time.Sleep(1 * time.Second)
											rpoip = 5
										}
									} else if b == 2 {
										if cursep == 0 {
											time.Sleep(1 * time.Second)
											fmt.Println("Sorry you don't have this item")
											time.Sleep(1 * time.Second)
											continue
										} else {
											cursep--
											time.Sleep(1 * time.Second)
											rcursep = 10
										}
									} else if b == 3 {
										if hpst == 0 {
											time.Sleep(1 * time.Second)
											fmt.Println("Sorry you don't have this item")
											time.Sleep(1 * time.Second)
											continue
										} else {
											hpst--
											if hp1 == hp1or {
												hp2 = hp2 - 30
												time.Sleep(1 * time.Second)
												fmt.Println("You're not low enough to heal")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took 30 damage")
											} else if hp1 >= (hp1or - 30){
												time.Sleep(1 * time.Second)
												hp1 = hp1or
												hp2 = hp2 - 30
												fmt.Println("You have been healed to max")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took 30 damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
											} else {
												time.Sleep(1 * time.Second)
												hp1 = hp1 + 30
												hp2 = hp2 - 30
												fmt.Println("You stole 30 hp")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
											}
										}
									}
								} else {
									fmt.Println("Sorry that is an invalid command")
								}
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(5)
								if random == 3 {
									if hp1 == hp1or / 2 {
										hp1 = hp1 + pethe
										time.Sleep(1 * time.Second)
										fmt.Println("You have been healed",pethe,"hp")
									} else {
										hp2 = hp2 - petat
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",petat,"damage")
									}
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}
								time.Sleep(1 * time.Second)
								fmt.Println("enemy is attacking")
								time.Sleep(1 * time.Second)
								if hp2 <= (hp2or / 3) && hp1 > attacken2 {
									time.Sleep(1 * time.Second)
									hp2 = hp2 + attacken3
									fmt.Println("enemy has been healed",attacken3,"health")
									time.Sleep(1 * time.Second)
									continue
								}
								if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp1 = hp1 - attacken1 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											//check = 2
											break
										}	
									}else {
										fmt.Println("you took",attacken1,"damage")
										hp1 = hp1-attacken1
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											//check = 2
											break
										}
									}
								} else {
									var b = time.Now().UnixNano()
									rand.Seed(b)
									var random3 = rand.Intn(5)
										if random3 == 1 || random3 == 2 || random3 == 3 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp1 = hp1 - attacken2 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													//check = 2
													break
												}	
											}else {
												fmt.Println("you took",attacken2,"damage")
												hp1 = hp1-attacken2
												time.Sleep(1 * time.Second)	
												if hp1 <= 0 {
													fmt.Println("You Died")
													//check = 2
													break
												}
											}
										} else {
											fmt.Println("attack missed")
												time.Sleep(1 * time.Second)	
												continue
										}
									}
									if rpoip != 0 {
										hp2 = hp2 - 10
										rpoip--
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy poisoned 10 damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									}
									if rcursep != 0 {
										if rcursep == 1 {
											hp2 = hp2 - 150
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took 150 damage.")
											time.Sleep(1 * time.Second)
											rcursep = 0
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}
										rcursep--
										time.Sleep(1 * time.Second)
										fmt.Println(rcursep,"turns left")
										time.Sleep(1 * time.Second)
									}
								}
								hp1 = hp1or
								time.Sleep(1 * time.Second)
								fmt.Println("You gained 50 coins")
								time.Sleep(1 * time.Second)
								mon = mon + 50
								fmt.Println("Total coins:",mon)
								time.Sleep(4 * time.Second)
								fmt.Println("The spider dies and he drops a spider eye, you can either use it for a weapon upgrade or magic, which do you choose?")
								time.Sleep(2 * time.Second)
								fmt.Println("1.Weapon")
								fmt.Println("2.Magic")
								fmt.Scanln(&b)
								time.Sleep(2 * time.Second)
								if b == 1 {
									attack1 = attack1 + 20
									attack2 = attack2 + 20
									attack4 = attack4 + 20
									fmt.Println("Weapon upgraded")
								} else {
									attack3 = attack3 + 10
									fmt.Println("Magic upgraded")
								}
								time.Sleep(2 * time.Second)
						}
				} else {
					fmt.Println("You climb up Mount Olympus to start your quest.")
					time.Sleep(4 * time.Second)
					fmt.Println("it's been two years since you visited this place, in those years Zeus was killed in a battle.")
					time.Sleep(4 * time.Second)
					fmt.Println("Now demons control Mount Olympus so you must find the demon that's in charge, and kill him.")
					time.Sleep(4 * time.Second)
					fmt.Println("You suddenly find yourself in front of a cursed minion.")
					time.Sleep(4 * time.Second)
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
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 3 {
									hp2 = hp2 - attack4 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}else {
									hp2 = hp2 - attack4
									fmt.Println("Enemy took",attack4,"damage")
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}
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
						fmt.Println()
						fmt.Println("choose your attack")
						fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
						fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
						fmt.Println("3.heal",attack3,"health 100% accuracy")
						fmt.Println("4.dark magic 100% accuracy")
						fmt.Scanln(&a)
						if a == 1 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(8)
							if random == 0 {
								hp2 = hp2 - attack1 * 3 / 2
								fmt.Println("Critical Hit")
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}	
							}else {
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy took",attack1,"damage")
								hp2 = hp2 - attack1
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}
							}
						} else if a == 2 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(3)
							if random == 1 || random == 2 || random == 3 {
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(8)
								if random == 3 {
									hp2 = hp2 - attack2 * 3 / 2
									fmt.Println("Critical Hit")
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
									time.Sleep(1 * time.Second)
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}else {
								time.Sleep(1 * time.Second)
								hp2 = hp2 - attack2
								fmt.Println("Enemy took",attack2,"damage")
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}
								}
							} else {
								time.Sleep(1 * time.Second)
								fmt.Println("attack missed")
								time.Sleep(1 * time.Second)
							}
						} else if a == 3 {
							if hp1 == hp1or {
								time.Sleep(1 * time.Second)
								fmt.Println("You're not low enough to heal")
								time.Sleep(1 * time.Second)
							} else if hp1 >= (hp1or - attack3){
								time.Sleep(1 * time.Second)
								hp1 = hp1or
								fmt.Println("You have been healed to max")
								time.Sleep(1 * time.Second)
							} else {
								time.Sleep(1 * time.Second)
								hp1 = hp1 + attack3
								fmt.Println("You have been healed",attack3,"health")
								time.Sleep(1 * time.Second)
							}
						} else if a == 4 {
							time.Sleep(1 * time.Second)
							fmt.Println("Which potion would you like to use")
							fmt.Println("1.Poison:",poip)
							fmt.Println("2.Curse",cursep)
							fmt.Println("3.Life Steal",hpst)
							fmt.Scanln(&b)
							time.Sleep(1 * time.Second)
							if b == 1 {
								if poip == 0 {
									time.Sleep(1 * time.Second)
									fmt.Println("Sorry you don't have this item")
									time.Sleep(1 * time.Second)
									continue
								} else {
									poip--
									time.Sleep(1 * time.Second)
									rpoip = 5
								}
							} else if b == 2 {
								if cursep == 0 {
									time.Sleep(1 * time.Second)
									fmt.Println("Sorry you don't have this item")
									time.Sleep(1 * time.Second)
									continue
								} else {
									cursep--
									time.Sleep(1 * time.Second)
									rcursep = 10
								}
							} else if b == 3 {
								if hpst == 0 {
									time.Sleep(1 * time.Second)
									fmt.Println("Sorry you don't have this item")
									time.Sleep(1 * time.Second)
									continue
								} else {
									hpst--
									if hp1 == hp1or {
										hp2 = hp2 - 30
										time.Sleep(1 * time.Second)
										fmt.Println("You're not low enough to heal")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took 30 damage")
									} else if hp1 >= (hp1or - 30){
										time.Sleep(1 * time.Second)
										hp1 = hp1or
										hp2 = hp2 - 30
										fmt.Println("You have been healed to max")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took 30 damage")
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
									} else {
										time.Sleep(1 * time.Second)
										hp1 = hp1 + 30
										hp2 = hp2 - 30
										fmt.Println("You stole 30 hp")
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
									}
								}
							}
						} else {
							fmt.Println("Sorry that is an invalid command")
						}
						var s = time.Now().UnixNano()
						rand.Seed(s)
						var random = rand.Intn(5)
						if random == 3 {
							if hp1 == hp1or / 2 {
								hp1 = hp1 + pethe
								time.Sleep(1 * time.Second)
								fmt.Println("You have been healed",pethe,"hp")
							} else {
								hp2 = hp2 - petat
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy took",petat,"damage")
							}
							if hp2 <= 0 {
								fmt.Println("You Win")
								break
							}	
						}
						time.Sleep(1 * time.Second)
						fmt.Println("enemy is attacking")
						time.Sleep(1 * time.Second)
						if hp2 <= (hp2or / 3) && hp1 > attacken2 {
							time.Sleep(1 * time.Second)
							hp2 = hp2 + attacken3
							fmt.Println("enemy has been healed",attacken3,"health")
							time.Sleep(1 * time.Second)
							continue
						}
						if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
							var s = time.Now().UnixNano()
							rand.Seed(s)
							var random = rand.Intn(8)
							if random == 0 {
								hp1 = hp1 - attacken1 * 3 / 2
								fmt.Println("Critical Hit")
								time.Sleep(1 * time.Second)
								fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
								time.Sleep(1 * time.Second)
								if hp1 <= 0 {
									fmt.Println("You Died")
									//check = 2
									break
								}	
							}else {
								fmt.Println("you took",attacken1,"damage")
								hp1 = hp1-attacken1
								time.Sleep(1 * time.Second)
								if hp1 <= 0 {
									fmt.Println("You Died")
									//check = 2
									break
								}
							}
						} else {
							var b = time.Now().UnixNano()
							rand.Seed(b)
							var random3 = rand.Intn(5)
								if random3 == 1 || random3 == 2 || random3 == 3 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp1 = hp1 - attacken2 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											//check = 2
											break
										}	
									}else {
										fmt.Println("you took",attacken2,"damage")
										hp1 = hp1-attacken2
										time.Sleep(1 * time.Second)	
										if hp1 <= 0 {
											fmt.Println("You Died")
											//check = 2
											break
										}
									}
								} else {
									fmt.Println("attack missed")
										time.Sleep(1 * time.Second)	
										continue
								}
							}
							if rpoip != 0 {
								hp2 = hp2 - 10
								rpoip--
								time.Sleep(1 * time.Second)
								fmt.Println("Enemy poisoned 10 damage")
								time.Sleep(1 * time.Second)
								if hp2 <= 0 {
									fmt.Println("You Win")
									break
								}	
							}
							if rcursep != 0 {
								if rcursep == 1 {
									hp2 = hp2 - 150
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy took 150 damage.")
									time.Sleep(1 * time.Second)
									rcursep = 0
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}
								rcursep--
								time.Sleep(1 * time.Second)
								fmt.Println(rcursep,"turns left")
								time.Sleep(1 * time.Second)
							}
						}
						hp1 = hp1or
						time.Sleep(1 * time.Second)
						fmt.Println("You gained 100 coins")
						time.Sleep(1 * time.Second)
						mon = mon + 100
						fmt.Println("Total coins:",mon)
						attack1 = attack1 + 20
						attack2 = attack2 + 20
						attack4 = attack4 + 20
						time.Sleep(4 * time.Second)
						fmt.Println("After effortlessly vanquishing the minion you take his spear and continue forwards.")
						time.Sleep(4 * time.Second)
						fmt.Println("Weapon upgrade")
						time.Sleep(4 * time.Second)
						fmt.Println("While you continue to travel up the cliff you see two routes.")
						time.Sleep(4 * time.Second)
						fmt.Println("You can either go to a small cave or you can keep climbing and go to the top of the mountain.")
						time.Sleep(4 * time.Second)
						fmt.Println("1.Cave")
						fmt.Println("2.Climb up")
						fmt.Scanln(&b)
						if b == 1 {
							fmt.Println("You walk into the dark cave and several cave spiders swarm you.")
							time.Sleep(4 * time.Second)
							fmt.Println("You try to fight them off but it's hopeless, You are then pushed out of the cave and fall into the jagged rocks.")
							break
						} else {
							fmt.Println("You climb up to the top and find yourself in front of a cursed monkey.")
							time.Sleep(4 * time.Second)
							rpoip = 0
							rcursep = 0
							hp2 = 100
							hp2or = hp2
							attacken1 = 10
							attacken2 = 30
							attacken3 = 15
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
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 3 {
											hp2 = hp2 - attack4 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack4 + (attack4 / 2),"damage")
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
											hp2 = hp2 - attack4
											fmt.Println("Enemy took",attack4,"damage")
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}
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
								fmt.Println()
								fmt.Println("choose your attack")
								fmt.Println("1.light attack ",attack1,"damage 100% accuracy")
								fmt.Println("2.heavy attack ",attack2,"damage 75% accuracy")
								fmt.Println("3.heal",attack3,"health 100% accuracy")
								fmt.Println("4.dark magic 100% accuracy")
								fmt.Scanln(&a)
								if a == 1 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp2 = hp2 - attack1 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack1 + (attack1 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									} else {
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",attack1,"damage")
										hp2 = hp2 - attack1
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
									}
								} else if a == 2 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(3)
									if random == 1 || random == 2 || random == 3 {
										var s = time.Now().UnixNano()
										rand.Seed(s)
										var random = rand.Intn(8)
										if random == 3 {
											hp2 = hp2 - attack2 * 3 / 2
											fmt.Println("Critical Hit")
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took",attack2 + (attack2 / 2),"damage")
											time.Sleep(1 * time.Second)
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}else {
										time.Sleep(1 * time.Second)
										hp2 = hp2 - attack2
										fmt.Println("Enemy took",attack2,"damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}
										}
									} else {
										time.Sleep(1 * time.Second)
										fmt.Println("attack missed")
										time.Sleep(1 * time.Second)
									}
								} else if a == 3 {
									if hp1 == hp1or {
										time.Sleep(1 * time.Second)
										fmt.Println("You're not low enough to heal")
										time.Sleep(1 * time.Second)
									} else if hp1 >= (hp1or - attack3){
										time.Sleep(1 * time.Second)
										hp1 = hp1or
										fmt.Println("You have been healed to max")
										time.Sleep(1 * time.Second)
									} else {
										time.Sleep(1 * time.Second)
										hp1 = hp1 + attack3
										fmt.Println("You have been healed",attack3,"health")
										time.Sleep(1 * time.Second)
									}
								} else if a == 4 {
									time.Sleep(1 * time.Second)
									fmt.Println("Which potion would you like to use")
									fmt.Println("1.Poison:",poip)
									fmt.Println("2.Curse",cursep)
									fmt.Println("3.Life Steal",hpst)
									fmt.Scanln(&b)
									time.Sleep(1 * time.Second)
									if b == 1 {
										if poip == 0 {
											time.Sleep(1 * time.Second)
											fmt.Println("Sorry you don't have this item")
											time.Sleep(1 * time.Second)
											continue
										} else {
											poip--
											time.Sleep(1 * time.Second)
											rpoip = 5
										}
									} else if b == 2 {
										if cursep == 0 {
											time.Sleep(1 * time.Second)
											fmt.Println("Sorry you don't have this item")
											time.Sleep(1 * time.Second)
											continue
										} else {
											cursep--
											time.Sleep(1 * time.Second)
											rcursep = 10
										}
									} else if b == 3 {
										if hpst == 0 {
											time.Sleep(1 * time.Second)
											fmt.Println("Sorry you don't have this item")
											time.Sleep(1 * time.Second)
											continue
										} else {
											hpst--
											if hp1 == hp1or {
												hp2 = hp2 - 30
												time.Sleep(1 * time.Second)
												fmt.Println("You're not low enough to heal")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took 30 damage")
											} else if hp1 >= (hp1or - 30){
												time.Sleep(1 * time.Second)
												hp1 = hp1or
												hp2 = hp2 - 30
												fmt.Println("You have been healed to max")
												time.Sleep(1 * time.Second)
												fmt.Println("Enemy took 30 damage")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
											} else {
												time.Sleep(1 * time.Second)
												hp1 = hp1 + 30
												hp2 = hp2 - 30
												fmt.Println("You stole 30 hp")
												if hp2 <= 0 {
													fmt.Println("You Win")
													break
												}
											}
										}
									}
								} else {
									fmt.Println("Sorry that is an invalid command")
								}
								var s = time.Now().UnixNano()
								rand.Seed(s)
								var random = rand.Intn(5)
								if random == 3 {
									if hp1 == hp1or / 2 {
										hp1 = hp1 + pethe
										time.Sleep(1 * time.Second)
										fmt.Println("You have been healed",pethe,"hp")
									} else {
										hp2 = hp2 - petat
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy took",petat,"damage")
									}
									if hp2 <= 0 {
										fmt.Println("You Win")
										break
									}	
								}
								time.Sleep(1 * time.Second)
								fmt.Println("enemy is attacking")
								time.Sleep(1 * time.Second)
								var e = time.Now().UnixNano()
								rand.Seed(e)
								var random69 = rand.Intn(5)
								if random69 == 2 {
									time.Sleep(1 * time.Second)
									fmt.Println("Enemy uses special attack: Monkey flip")
									time.Sleep(1 * time.Second)
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp1 = hp1 - 40 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("you took",40 + (40 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											//check = 2
											break
										}	
									} else {
										fmt.Println("you took 40 damage")
										hp1 = hp1 - 40
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											//check = 2
											break
										}
									}
									continue
								}
								if hp2 <= (hp2or / 3) && hp1 > attacken2 {
									time.Sleep(1 * time.Second)
									hp2 = hp2 + attacken3
									fmt.Println("enemy has been healed",attacken3,"health")
									time.Sleep(1 * time.Second)
									continue
								}
								if hp1 <= attacken1 || hp1 <= attacken1 * 2 && hp2 > hp1 {
									var s = time.Now().UnixNano()
									rand.Seed(s)
									var random = rand.Intn(8)
									if random == 0 {
										hp1 = hp1 - attacken1 * 3 / 2
										fmt.Println("Critical Hit")
										time.Sleep(1 * time.Second)
										fmt.Println("you took",attacken1 + (attacken1 / 2),"damage")
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											//check = 2
											break
										}	
									} else {
										fmt.Println("you took",attacken1,"damage")
										hp1 = hp1-attacken1
										time.Sleep(1 * time.Second)
										if hp1 <= 0 {
											fmt.Println("You Died")
											//check = 2
											break
										}
									}
								} else {
									var b = time.Now().UnixNano()
									rand.Seed(b)
									var random3 = rand.Intn(5)
										if random3 == 1 || random3 == 2 || random3 == 3 {
											var s = time.Now().UnixNano()
											rand.Seed(s)
											var random = rand.Intn(8)
											if random == 0 {
												hp1 = hp1 - attacken2 * 3 / 2
												fmt.Println("Critical Hit")
												time.Sleep(1 * time.Second)
												fmt.Println("you took",attacken2 + (attacken2 / 2),"damage")
												time.Sleep(1 * time.Second)
												if hp1 <= 0 {
													fmt.Println("You Died")
													//check = 2
													break
												}	
											}else {
												fmt.Println("you took",attacken2,"damage")
												hp1 = hp1-attacken2
												time.Sleep(1 * time.Second)	
												if hp1 <= 0 {
													fmt.Println("You Died")
													//check = 2
													break
												}
											}
										} else {
											fmt.Println("attack missed")
												time.Sleep(1 * time.Second)	
												continue
										}
									}
									if rpoip != 0 {
										hp2 = hp2 - 10
										rpoip--
										time.Sleep(1 * time.Second)
										fmt.Println("Enemy poisoned 10 damage")
										time.Sleep(1 * time.Second)
										if hp2 <= 0 {
											fmt.Println("You Win")
											break
										}	
									}
									if rcursep != 0 {
										if rcursep == 1 {
											hp2 = hp2 - 150
											time.Sleep(1 * time.Second)
											fmt.Println("Enemy took 150 damage.")
											time.Sleep(1 * time.Second)
											rcursep = 0
											if hp2 <= 0 {
												fmt.Println("You Win")
												break
											}	
										}
										rcursep--
										time.Sleep(1 * time.Second)
										fmt.Println(rcursep,"turns left")
										time.Sleep(1 * time.Second)
									}
								}
								hp1 = hp1or
								time.Sleep(1 * time.Second)
								fmt.Println("You gained 69 coins")
								time.Sleep(1 * time.Second)
								mon = mon + 69
								fmt.Println("Total coins:",mon)
								time.Sleep(4 * time.Second)
								fmt.Println("The monkey dies and you get his hide, you can either upgrade your magic or armor.")
								time.Sleep(2 * time.Second)
								fmt.Println("1.Armor")
								fmt.Println("2.Magics")
								fmt.Scanln(&b)
						}
				}
			} else {
				var c int
				for {
					fmt.Println("1.How to play Campaign")
					fmt.Println("2.How to play PVP")
					fmt.Println("3.How to play Playground")
					fmt.Println("4.How to play Survival")
					fmt.Println("5.Back")
					fmt.Scanln(&c)
					time.Sleep(1 * time.Second)
					if c == 1 {
						fmt.Println("The game will ask you questions and you get to choose what to do based on the choices.")
						fmt.Println("In campaign mode your choices affect the results. You can level up and fight bosses to complete the game.")
					} else if c == 2 {
						fmt.Println("The PVP mode allows you to play against your friends, Before each PVP match you have time to buy supplies and prepare to face your friend.")
					} else if c == 3 {
						fmt.Println("Playground mode gives you infinite coins, health, and damage.")
					}else if c == 4 {
						fmt.Println("Survival is where you fight a series of enemys until you die. Each round the enemy gets stronger.")
					} else {
						break
					}
					time.Sleep(5 * time.Second)
				}
			}
		}
	}