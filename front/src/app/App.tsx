import React from 'react';
import styles from './App.module.scss';
import {AppRouter} from './AppRouter';
import {LoadableContent} from '../components/LoadableContent';
import {useInit} from './hooks';

export type AppProps = Record<string, never>;

export const App: React.FC<AppProps> = () => {
    const {pending} = useInit();
    return (
        <div className={styles.root}>
            <LoadableContent isLoading={pending}>
                <AppRouter />
            </LoadableContent>
        </div>
    );
};
