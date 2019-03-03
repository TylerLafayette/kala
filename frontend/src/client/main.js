import React, { Component } from "react"
import { render } from "react-dom"
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom"

import Home from "./pages/Home.jsx"
import ControlPanel from "./pages/ControlPanel.jsx"

const API_URL = "http://10.178.83.26:7800/api"

class Main extends Component {
    constructor() {
        super()
        this.state = {
            loading: true,
            data: {},
            amountSpent: 0,
            transactions: [],
            suggestions: []
        }
        this.interval = null
    }
    async componentDidMount() {
        let response = await fetch(API_URL + "/overview")
        let json = await response.json()
        this.setState({
            loading: false,
            amountSpent: json.amountSpent,
            data: json,
            transactions: json.transactions,
            suggestions: json.suggestions
        })
        this.interval = setInterval(this.update.bind(this), 1000)
    }
    componentWillUnmount() {
        clearInterval(this.interval)
    }
    // This is a terrible way to update in real-time.
    async update() {
        let response = await fetch(API_URL + "/overview")
        let json = await response.json()
        let amountSpent = this.state.amountSpent
        if(json.transactions.length > this.state.transactions.length) amountSpent += this.state.transactions[0].price
        if(json.transactions.length < this.state.transactions.length) amountSpent = 500

        if(json.transactions.length != this.state.transactions.length || json.suggestions.length != this.state.suggestions.length)
            this.setState({
                loading: false,
                data: json,
                amountSpent,
                transactions: json.transactions,
                suggestions: json.suggestions,
            })
    }
    render() {
        if(this.state.loading)
            return (
                <div style={{
                    display: "flex",
                    width: "100%",
                    height: "100%",
                    alignItems: "center",
                    justifyContent: "center"
                }}>
                    <div style={{
                        width: "40px",
                        height: "40px",
                        borderRadius: "30px",
                        border: "2px dotted #000000",
                        animation: "spinny-circle 3s linear infinite"
                    }}></div>
                </div>
            )

        return (
            <div className="fade-in">
                <Router>
                    <Switch>
                        <Route exact path="/" component={() => <Home amountSpent={this.state.amountSpent} data={this.state.data} transactions={this.state.transactions} suggestions={this.state.suggestions} />} />
                        <Route exact path="/ctrl" component={() => <ControlPanel data={this.state.data} />} />
                    </Switch>
                </Router>
            </div>
        )
    }
}

render(<Main />, document.getElementById("app-mount"))