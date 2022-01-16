import {FieldInputProps} from 'formik';

export enum FieldType {
    Select = 'select',
    Input = 'input',
    Checkbox = 'checkbox',
    TextArea = 'textarea',
}

export type FieldOption = {
    value: string;
    title: string;
};

export type SetFieldValue = (field: string, value: unknown) => void;

export type BaseFieldProps = FieldInputProps<
    string | ReadonlyArray<string> | number | undefined
> & {
    error?: string | string[];
    id: string;
    title: string;
    type: FieldType | string;
    required: boolean;
    setFieldValue: SetFieldValue;
};

export type SelectFieldProps = BaseFieldProps & {
    options: FieldOption[];
    type: 'select';
};

export type CheckboxFieldProps = BaseFieldProps & {
    options: FieldOption[];
    type: 'checkbox';
};

export type TextAreaFieldProps = BaseFieldProps;

export type NumberFieldProps = BaseFieldProps & {
    min?: number;
    max?: number;
    type: 'number';
};

export type FieldProps =
    | BaseFieldProps
    | SelectFieldProps
    | NumberFieldProps
    | CheckboxFieldProps
    | TextAreaFieldProps;

export type DefaultFiledItem<T> = Pick<FieldProps, 'title' | 'required' | 'type'> & {
    id: keyof T;
};

export type SelectFieldItem<T> = DefaultFiledItem<T> &
    Pick<SelectFieldProps, 'options' | 'type'> & {
        type: 'select';
    };

export type NumberFieldItem<T> = DefaultFiledItem<T> &
    Pick<NumberFieldProps, 'min' | 'max' | 'type'>;

export type CheckboxFieldItem<T> = DefaultFiledItem<T> &
    Pick<CheckboxFieldProps, 'options' | 'type'>;

export type TextAreaFieldItem<T> = DefaultFiledItem<T>;

export type FieldItem<T> =
    | DefaultFiledItem<T>
    | SelectFieldItem<T>
    | NumberFieldItem<T>
    | CheckboxFieldItem<T>
    | TextAreaFieldItem<T>;
