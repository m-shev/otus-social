import React from 'react';
import styles from './Field.module.scss';
import {BaseFieldProps, NumberField, SelectField} from './types';

export type FieldSelectOption = {
    value: string;
    title: string;
};

const isSelectedField = (props: FieldProps): props is SelectField => {
    return props.type === 'select';
};

const isDefaultField = (props: FieldProps): props is SelectField => {
    return props.type !== 'select';
};

export type FieldProps = BaseFieldProps | SelectField | NumberField;

export const Field: React.FC<FieldProps> = (props) => {
    return (
        <div className={styles.root}>
            <label htmlFor={props.id} className={styles.label}>
                {props.title}
            </label>

            {isSelectedField(props) && (
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

            {isDefaultField(props) && <input className={styles.field} {...props} />}
        </div>
    );
};
