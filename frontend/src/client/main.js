import React, { Component } from "react"
import { render } from "react-dom"

import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom"

class Main extends Component {
    render() {
        return (
            <h1>Hello world!</h1>
        )
    }
}

render(<Main />, document.getElementById("app-mount"))