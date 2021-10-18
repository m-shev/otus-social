import React from 'react';
import styles from './Field.module.scss';
import {FieldProps} from './types';
import {SpecificField} from './SpecificField';

export const Field: React.FC<FieldProps> = (props) => {
    return (
        <div className={styles.root}>
            <label htmlFor={props.id} className={styles.label}>
                {props.title}
            </label>

            <SpecificField {...props} />
        </div>
    );
};
