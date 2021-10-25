import * as React from 'react';
import {BrowserRouter, Route, Switch} from 'react-router-dom';
import {RegistrationPage} from '../pages/RegistrationPage';
import {MainPage} from '../pages/MainPage';
import {LoginPage} from '../pages/LoginPage';
import {ProfilePage} from '../pages/ProfilePage';
import {FindUserPage} from '../pages/FindUserPage';

export const AppRouter: React.FC = () => {
    return (
        <BrowserRouter>
            <Switch>
                <Route exact path={'/registration'} component={RegistrationPage} />
                <Route exact path={'/login'} component={LoginPage} />
                <Route exact path={'/profile'} component={ProfilePage} />
                <Route exact path={'/find'} component={FindUserPage} />
                <Route exact path={'/'} component={MainPage} />
            </Switch>
        </BrowserRouter>
    );
};
