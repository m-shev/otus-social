import * as React from 'react';
import {useDots} from './useDots';
import styles from './styles.module.scss';

export const DefaultLoader: React.FC = () => {
    const dots = useDots();
    return (
        <div className={styles.root}>
            <div>Выполняю загрузку данных</div>
            <div className={styles.dots}>{dots}</div>
        </div>
    );
};
