package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type cardType int

const (
	Jocker cardLabel = iota + 1
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

const (
	HighCard cardType = iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type cardLabel int

type Card struct {
	rank     int
	bid      int
	hand     []cardLabel
	cardType cardType
}

func parse1(input string) []Card {
	lines := strings.Split(input, "\n")
	cards := make([]Card, 0, len(lines))
	for _, line := range lines {
		handBid := strings.Split(line, " ")
		bid, err := strconv.Atoi(handBid[1])
		if err != nil {
			log.Fatalf("Wrong bid: %s", handBid[1])
		}

		hand := make([]cardLabel, 0, len(handBid[0]))
		handMap := make(map[cardLabel]int, len(handBid[0]))

		for _, ch := range handBid[0] {
			switch ch {
			case 'T':
				hand = append(hand, Ten)
				handMap[Ten]++
			case 'J':
				hand = append(hand, Jack)
				handMap[Jack]++
			case 'Q':
				hand = append(hand, Queen)
				handMap[Queen]++
			case 'K':
				hand = append(hand, King)
				handMap[King]++
			case 'A':
				hand = append(hand, Ace)
				handMap[Ace]++
			default:
				num := cardLabel(ch - '0')
				if num < 2 || num > 9 {
					log.Fatalf("Wrong card in hand: %v", ch)
				}
				hand = append(hand, num)
				handMap[num]++
			}
		}

		values := make([]string, 0, len(handMap))
		for _, v := range handMap {
			values = append(values, strconv.Itoa(v))
		}
		slices.Sort(values)
		typeStr := strings.Join(values, "")

		var cardType cardType
		switch typeStr {
		case "5":
			cardType = FiveOfAKind
		case "14":
			cardType = FourOfAKind
		case "23":
			cardType = FullHouse
		case "113":
			cardType = ThreeOfAKind
		case "122":
			cardType = TwoPair
		case "1112":
			cardType = OnePair
		case "11111":
			cardType = HighCard
		default:
			log.Fatalf("Wrong type card: %s", typeStr)
		}

		card := Card{
			bid:      bid,
			hand:     hand,
			cardType: cardType,
		}
		cards = append(cards, card)
	}
	return cards
}

func getCardTypeStr(handMap map[cardLabel]int, card cardLabel) string {
	jockerCount, contains := handMap[card]
	if contains {
		if jockerCount == 5 {
			return "5"
		}
		values := make([]int, 0, len(handMap))
		for k, v := range handMap {
			if k == card {
				continue
			}
			values = append(values, v)
		}
		slices.Sort(values)
		values[len(values)-1] += jockerCount
		valuesStr := make([]string, 0, len(values))
		for _, v := range values {
			valuesStr = append(valuesStr, strconv.Itoa(v))
		}
		return strings.Join(valuesStr, "")
	}
	values := make([]string, 0, len(handMap))
	for _, v := range handMap {
		values = append(values, strconv.Itoa(v))
	}
	slices.Sort(values)
	return strings.Join(values, "")
}

func parse2(input string) []Card {
	lines := strings.Split(input, "\n")
	cards := make([]Card, 0, len(lines))
	for _, line := range lines {
		handBid := strings.Split(line, " ")
		bid, err := strconv.Atoi(handBid[1])
		if err != nil {
			log.Fatalf("Wrong bid: %s", handBid[1])
		}

		hand := make([]cardLabel, 0, len(handBid[0]))
		handMap := make(map[cardLabel]int, len(handBid[0]))

		for _, ch := range handBid[0] {
			switch ch {
			case 'T':
				hand = append(hand, Ten)
				handMap[Ten]++
			case 'J':
				hand = append(hand, Jocker)
				handMap[Jocker]++
			case 'Q':
				hand = append(hand, Queen)
				handMap[Queen]++
			case 'K':
				hand = append(hand, King)
				handMap[King]++
			case 'A':
				hand = append(hand, Ace)
				handMap[Ace]++
			default:
				num := cardLabel(ch - '0')
				if num < 2 || num > 9 {
					log.Fatalf("Wrong card in hand: %v", ch)
				}
				hand = append(hand, num)
				handMap[num]++
			}
		}
		typeStr := getCardTypeStr(handMap, Jocker)

		var cardType cardType
		switch typeStr {
		case "5":
			cardType = FiveOfAKind
		case "14":
			cardType = FourOfAKind
		case "23":
			cardType = FullHouse
		case "113":
			cardType = ThreeOfAKind
		case "122":
			cardType = TwoPair
		case "1112":
			cardType = OnePair
		case "11111":
			cardType = HighCard
		default:
			log.Fatalf("Wrong type card: %s", typeStr)
		}

		card := Card{
			bid:      bid,
			hand:     hand,
			cardType: cardType,
		}
		cards = append(cards, card)
	}
	return cards
}

func getRes(cards []Card) int {
	sort.Slice(cards, func(i, j int) bool {
		if cards[i].cardType != cards[j].cardType {
			return cards[i].cardType < cards[j].cardType
		}
		var res bool
		switch slices.Compare(cards[i].hand, cards[j].hand) {
		case -1:
			res = true
		case 1:
			res = false
		default:
			log.Fatalf("Unreachable")
		}
		return res
	})
	var total int
	for i, card := range cards {
		card.rank = i + 1
		total += card.rank * card.bid
	}
	return total
}

func part1(input string) int {
	cards := parse1(input)
	return getRes(cards)
}

func part2(input string) int {
	cards := parse2(input)
	return getRes(cards)
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Not file")
	}
	input := string(data)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
