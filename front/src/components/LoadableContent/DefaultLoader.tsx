import * as React from 'react';
import {useDots} from './useDots';
import styles from './styles.module.scss';

export type DefaultLoaderProps = {};

export const DefaultLoader: React.FC<DefaultLoaderProps> = () => {
    const dots = useDots();
    return (
        <div className={styles.root}>
            <div>Выполняю загрузку данных</div>
            <div className={styles.dots}>{dots}</div>
        </div>
    );
};
