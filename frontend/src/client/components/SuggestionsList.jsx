import React, { Component } from "react"

import "./SuggestionsList.scss"

export default class SuggestionsList extends Component {
    constructor() {
        super()
        this.state = {
            collapsed: []
        }
        this.renderButtons = this.renderButtons.bind(this)
    }
    renderButtons(i) {
        const { callToActionType: type, data } = i
        if(type == "SAVING_SCHEDULE")
            return (
                <a href="#" className="cta-button">
                    Setup ${data.amtMonthly}/mo schedule
                </a>
            )
        if(type == "LINK_BUTTON")
            return (
                <a href={data.url} target="_blank" className="cta-button">
                    {data.text}
                </a>
            )
    }
    collapse(index) {
        this.setState({
            collapsed: [...this.state.collapsed, index]
        })
    }
    render() {
        const { suggestions } = this.props

        return (
            <div className={`suggestions-list animated`}>
                {suggestions.map((item, i) => (
                    <div className={`suggestion-row ${this.state.collapsed.includes(i) ? "collapsed": ""}`} style={{
                        animationDelay: (i * 0.1) + "s"
                    }}>
                        <div className="row">
                            <span className="emoji">{ item.emoji }</span>
                            <div className="title-and-description">
                                <span className="header">{ item.header }</span>
                                <span className="description">{ item.description.replace("{}", item.amountSpentThisMonth) }</span>
                            </div>
                        </div>
                        <div className="row">
                            {
                                this.renderButtons(item)
                            }
                            <a href="#" onClick={this.collapse.bind(this, i)} className="cta-button danger">
                                Dismiss
                            </a>
                        </div>
                    </div>
                ))}
            </div>
        )
    }
}