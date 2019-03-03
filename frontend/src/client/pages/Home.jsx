import React, { Component } from "react"

import "./Home.scss"
import Navbar from "../components/Navbar"
import BudgetMeter from "../components/BudgetMeter"
import TransactionsList from "../components/TransactionsList"
import SuggestionsList from "../components/SuggestionsList"

export default class Home extends Component {
    render() {
        const { data, transactions, suggestions } = this.props

        return (
            <div className="home-page-wrapper">
                <Navbar />
                <span className="hello-span animated">Welcome back, Lucas.</span>
                <BudgetMeter className="animated" amountSpent={data.amountSpent} budget={data.budget} />
                <div className="flex-split">
                    <TransactionsList transactions={transactions} />
                    <SuggestionsList suggestions={suggestions} />
                </div>
            </div>
        )
    }
}