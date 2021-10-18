import * as React from 'react';
import {Link} from 'react-router-dom';
import styles from './MainPage.module.scss';

export type MainPageProps = {};

export const MainPage: React.FC<MainPageProps> = () => {
    return (
        <div className={styles.root}>
            <h1>Социальная сеть</h1>

            <div className={styles.linkWrapper}>
                <Link to="/login">Войти</Link>
                <Link to="/registration">Регистрация</Link>
            </div>
        </div>
    );
};
