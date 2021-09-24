import React from 'react';
import styles from './Field.module.scss';
import {BaseFieldProps, NumberFieldProps, SelectFieldProps} from './types';

export type FieldProps = BaseFieldProps | SelectFieldProps | NumberFieldProps;

const isSelectedFieldProps = (props: FieldProps): props is SelectFieldProps => {
    return props.type === 'select';
};

const isBaseFieldProps = (props: FieldProps): props is BaseFieldProps | NumberFieldProps => {
    return props.type !== 'select';
};

export const Field: React.FC<FieldProps> = (props) => {
    return (
        <div className={styles.root}>
            <label htmlFor={props.id} className={styles.label}>
                {props.title}
            </label>

            {isSelectedFieldProps(props) && (
                <select
                    id={props.id}
                    className={styles.field}
                    value={props.value}
                    onChange={props.onChange}
                >
                    {props.options.map((option) => {
                        return (
                            <option value={option.value} key={option.value}>
                                {option.title}
                            </option>
                        );
                    })}
                </select>
            )}

            {isBaseFieldProps(props) && <input className={styles.field} {...props} />}
        </div>
    );
};
