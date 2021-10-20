import * as React from 'react';
import {CheckboxFieldProps, FieldProps, FieldType, SelectFieldProps} from './types';
import styles from './Field.module.scss';
import {CheckboxGroup} from './CheckboxGroup';

export type SpecificFieldProps = FieldProps;

const isSelectedFieldProps = (props: FieldProps): props is SelectFieldProps => {
    return props.type === FieldType.Select;
};

const isCheckboxProps = (props: FieldProps): props is CheckboxFieldProps => {
    return props.type === FieldType.Checkbox;
};

export const SpecificField: React.FC<SpecificFieldProps> = (props) => {
    if (isSelectedFieldProps(props)) {
        return (
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
        );
    }

    if (isCheckboxProps(props)) {
        return <CheckboxGroup {...props} />;
    }

    return <input className={styles.field} {...props} />;
};
