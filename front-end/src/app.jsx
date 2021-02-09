import Home from './components/home.jsx';
import ErrorPage from './components/ErrorPage.jsx';
import WeaponsPage from './components/weapons.jsx';
import { BrowserRouter, Switch, Route } from 'react-router-dom';

const App = () => {
    return (
        <BrowserRouter>
            <Route exact path='/' component={WeaponsPage} />
        </BrowserRouter>
    )
}

export default App;
