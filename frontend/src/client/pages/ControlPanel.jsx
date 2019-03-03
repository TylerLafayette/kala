import React, { Component } from "react"

import "./ControlPanel.scss"

const API_URL = "http://10.178.83.26:7800/api"

export default class ControlPanel extends Component {
    async _send() {
        let description = this.refs.description.value
        let price = parseFloat(this.refs.price.value)
        let final_balance = parseFloat(this.refs.price.value)
        let response = await fetch(API_URL + "/insert", {
            method: "POST",
            headers: {
                "Accept": "application/json",
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                description,
                price,
                final_balance
            })
        })
        let json = await response.json()
        alert(JSON.stringify(json))
    }
    async _reset() {
        let response = await fetch(API_URL + "/loadTransactionsIntoDb")
        let json = await response.json()
        alert(JSON.stringify(json))
    }
    render() {
        return (
            <div className="control-panel">
                <input type="text" className="input" ref="description" placeholder="Description" />
                <input type="number" className="input" ref="price" placeholder="Price" />
                <input type="number" className="input" ref="newBal" placeholder="New balance" />
                <input type="submit" onClick={this._send.bind(this)} className="submit" value="Send" />
                <input type="submit" onClick={this._reset.bind(this)} className="submit danger" value="Reset" />
            </div>
        )
    }
}