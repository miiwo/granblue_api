import React, { Component } from 'react';
import SearchBar from './searchbar.jsx';

function WeaponDetails(props) {
    return (
        <>
            <label htmlFor='wep_name' className='font-weight-bold user-select-none'>Name:</label>
            <p id='wep_name' className='pb-3'>{props.weapon.name}</p>
            <label htmlFor='wep_des' className='font-weight-bold user-select-none'>Description:</label>
            <p id='wep_des'>This is a description.</p>
            <a href='http://placehold.it/' target='_blank' className='thumbnail'><img src='http://placehold.it/200' alt='...' className='img-fluid img-thumbnail' /></a>
        </>
    )
}

class WeaponsPage extends Component {

    constructor(props) {
        super(props);
        this.state = {weapons: [
            {name: 'Kappa Shoot'},
            {name: 'Flower Parasol'},
            {name: 'Kappa Shoot'},
            {name: 'Trappa Shoot'},
            {name: 'Doki Doki Shoot'},
            {name: 'Kappa Shoot'},
            {name: 'Frappa Shoot'},
            {name: 'Slappa Shoot'},
            {name: 'Kappa Shoot'},
            {name: 'Kappa Shoot'},
            {name: 'Kappa Shoot'},
            {name: 'Kappa Shoot'},
        ]};
        
        this.onSearchClick=this.onSearchClick.bind(this);
    }

    componentDidMount() {
        //fetch('/weapons/all').then(res => res.json()).then(weapons => this.setState({ weapons }));
    }

    onSearchClick(e) {
        e.preventDefault();
        console.log(e.target[0].value); // The first element in the form is the input I'm looking for
        fetch('/weapons/' + e.target[0].value).then(res => res.json())
            .then(weapons => this.setState({ weapons }))
            .catch(err => console.log('An error has occured with fetching the backend data'));
    }

    // Notes: I needed to set height of parent to an actual value for it to work. 0 does not : (
    render () {
        return (
            <>
                <div className='sticky-top bg-dark pl-3'>
                    <h1>Weapons</h1>
                    <SearchBar action={this.onSearchClick} />
                </div>
                <section className='container-fluid'>
                    {this.state.weapons.map((weapon, i) =>
                        <div key={i} className='rounded float-left border border-light m-2 pt-1 pb-3 pl-3 pr-3'>
                            <WeaponDetails weapon={weapon} />
                        </div>
                    )}
                </section>
            </>
        )
    }
}

export default WeaponsPage;
