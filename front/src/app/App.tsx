import React from 'react';
import styles from './App.module.scss';
import {AppRouter} from './AppRouter';

export type AppProps = Record<string, never>;

export const App: React.FC<AppProps> = () => {
    return (
        <div className={styles.root}>
            <AppRouter />
        </div>
    );
};
