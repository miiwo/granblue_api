import React, { Component } from 'react';
import SearchBar from './searchbar.jsx';

function WeaponDetails(props) {
    return (
        <>
            <p>{props.weapon.name}</p>
            <p>This is a description.</p>
            <a href='#' className='thumbnail'><img src='...' alt='...' /></a>
        </>
    )
}

class WeaponsPage extends Component {

    constructor(props) {
        super(props);
        this.state = {weapons: [{name: 'Kappa Shoot'}, {name: 'Flower Parasol'} ]};
        this.onSearchClick=this.onSearchClick.bind(this);
    }

    componentDidMount() {
        //fetch('/weapons/all').then(res => res.json()).then(weapons => this.setState({ weapons }));
    }

    onSearchClick(e) {
        e.preventDefault();
        console.log(e.target.value);
        fetch('/weapons/all').then(res => res.json())
            .then(weapons => this.setState({ weapons }))
            .catch(err => console.log('An error has occured with fetching the backend data'));
    }

    render () {
        return (
            <div>
                <div className='page-header'><h1>Weapons</h1></div>
                <SearchBar action={this.onSearchClick} />
                <div className='d-inline-flex'>
                    {this.state.weapons.map((weapon, i) =>
                        <div key={i} className='container rounded border border-light m-2'>
                            <WeaponDetails weapon={weapon} />
                        </div>
                    )}
                </div>
            </div>
        )
    }
}

export default WeaponsPage;
