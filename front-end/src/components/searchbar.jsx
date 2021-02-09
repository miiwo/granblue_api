import React from 'react';

function SearchBar(props) {
    return (
        <form className='form-inline'>
            <div className='form-group'>
                <label htmlFor='searchform' className='mr-2 pl-2'>Search</label>
                <input type='text' id='searchform' placeholder='Type your search' />
                <button onClick={props.action} className='btn btn-dark m-3'>Go</button>
            </div>
        </form>
    )
}

export default SearchBar;
