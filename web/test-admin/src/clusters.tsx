import { Create, SimpleForm, TextInput, List, Datagrid, TextField, ReferenceField, Edit, EditButton, DateInput, ReferenceManyField, DateField, required } from "react-admin";

export const ClusterList = () => (
    <List>
        <Datagrid>
            <TextField source="id" />
            <TextField source="name" />
            <EditButton />
        </Datagrid>
    </List>
);

export const ClusterCreate = () => (
    <Create>
        <SimpleForm>
            <TextInput source="name" validate={[required()]} fullWidth />
        </SimpleForm>
    </Create>
);

export const ClusterEdit = () => (
    <Edit>
        <SimpleForm>
            <TextInput disabled label="Id" source="id" />
            <TextInput source="name" validate={required()} />
        </SimpleForm>
    </Edit>
);