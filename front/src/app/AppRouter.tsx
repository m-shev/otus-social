import * as React from 'react';
import {BrowserRouter, Route, Switch} from 'react-router-dom';
import {RegistrationPage} from '../pages/RegistrationPage';
import {MainPage} from '../pages/MainPage';

export type RouterProps = {};

export const AppRouter: React.FC<RouterProps> = () => {
    return (
        <BrowserRouter>
            <Switch>
                <Route exact path={'/registration'} component={RegistrationPage} />
                <Route exact path={'/'} component={MainPage} />
            </Switch>
        </BrowserRouter>
    );
};
