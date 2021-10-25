import * as React from 'react';
import {Link} from 'react-router-dom';
import styles from './MainPage.module.scss';
import {Header} from '../../components/Header';
import {useRedirectToProfile} from '../../hooks';

export const MainPage: React.FC = () => {
    useRedirectToProfile();

    return (
        <>
            <div className={styles.root}>
                <Header showLoginButton={false} />

                <div className={styles.linkWrapper}>
                    <Link to="/login">Войти</Link>
                    <Link to="/registration">Регистрация</Link>
                </div>
            </div>
        </>
    );
};
