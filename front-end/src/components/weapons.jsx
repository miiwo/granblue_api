import React, { Component } from 'react';

class WeaponsPage extends Component {

    constructor(props) {
        super(props);
        this.state = {weapons: []};
    }

    componentDidMount() {
        fetch('/weapons/all').then(res => res.json()).then(weapons => this.setState({ weapons }));
    }

    render () {
        return (
            <div>
                <h1>Weapons</h1>
                {this.state.weapons.map((weapon, i) =>
                    <div key={i}>{weapon.name}</div>
                )}
            </div>
        )
    }
}

export default WeaponsPage;
