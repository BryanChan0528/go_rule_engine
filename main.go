package main

import (
	"fmt"
	"test_rule_engine/rule_engine"

	"github.com/hyperjumptech/grule-rule-engine/logger"
)

// can be part of user service and a separate directory
type User struct {
	Name              string  `json:"name"`
	Username          string  `json:"username"`
	Email             string  `json:"email"`
	Age               int     `json:"age"`
	Gender            string  `json:"gender"`
	TotalOrders       int     `json:"total_orders"`
	AverageOrderValue float64 `json:"average_order_value"`
}

// can be moved to offer directory
type OfferService interface {
	CheckOfferForUser(user User) bool
}

type OfferServiceClient struct {
	ruleEngineSvc *rule_engine.RuleEngineSvc
}

func NewOfferService(ruleEngineSvc *rule_engine.RuleEngineSvc) OfferService {
	return &OfferServiceClient{
		ruleEngineSvc: ruleEngineSvc,
	}
}

func (svc OfferServiceClient) CheckOfferForUser(user User) bool {
	offerCard := rule_engine.NewUserOfferContext()
	offerCard.UserOfferInput = &rule_engine.UserOfferInput{
		Name:              user.Name,
		Username:          user.Username,
		Email:             user.Email,
		Gender:            user.Gender,
		Age:               user.Age,
		TotalOrders:       user.TotalOrders,
		AverageOrderValue: user.AverageOrderValue,
	}

	err := svc.ruleEngineSvc.Execute(offerCard)
	if err != nil {
		logger.Log.Error("get user offer rule engine failed", err)
	}

	return offerCard.UserOfferOutput.IsOfferApplicable
}

func main() {
	ruleEngineSvc := rule_engine.NewRuleEngineSvc()
	offerSvc := NewOfferService(ruleEngineSvc)

	userA := User{
		Name:              "User X",
		Username:          "User1212",
		Email:             "user1212@domain.com",
		Gender:            "Male",
		Age:               20,
		TotalOrders:       30,
		AverageOrderValue: 225,
	}

	fmt.Println("offer validity for user A: ", offerSvc.CheckOfferForUser(userA))

	userB := User{
		Name:              "User Y",
		Username:          "User2323",
		Email:             "user2323@domain.com",
		Gender:            "Male",
		Age:               22,
		TotalOrders:       10,
		AverageOrderValue: 80,
	}

	fmt.Println("offer validity for user B: ", offerSvc.CheckOfferForUser(userB))
}
