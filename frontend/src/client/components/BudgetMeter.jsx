import React, { Component } from "react"

import "./BudgetMeter.scss"

export default class BudgetMeter extends Component {
    render() {
        return (
            <div className={`meter-wrapper ${this.props.className}`}>
                <span className="title-span">You've spent</span>
                <span className="amount-spent">${this.props.amountSpent}<span className="small">/${this.props.budget}</span></span>
                <div className="meter-outer">
                    <div className="meter-inner" style={{
                        width: this.props.amountSpent*100/this.props.budget + "%"
                    }} />
                </div>
            </div>
        )
    }
}