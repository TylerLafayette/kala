import React, { Component } from "react"

import "./TransactionsList.scss"

export default class TransactionsList extends Component { 
    truncateString(str, num) {
        if (str.length > num)
            return str.slice(0, num) + "..."
        return str
    }
    render() {
        const { transactions } = this.props

        const colors = {
            "coffee": "#593C1F",
            "restaurants": "#C242F4",
            "water": "#0cdaff",
            "groceries": "#82ff51",
            "snacks": "#FF3E88"
        }

        return (
            <div className="transactions-list animated">
                {transactions.map((item, i) => (
                    <div className="transaction-row" style={{
                        animationDelay: (i * 0.1) + "s"
                    }}>
                        <span className="description">{item.category.length > 0 ? <span style={{
                            backgroundColor: colors[item.category]
                        }} className="dot"></span> : null}{ this.truncateString(item.description, 26) }</span>
                        <span className="price">${ item.price }</span>
                    </div>
                ))}
            </div>
        )
    }
}