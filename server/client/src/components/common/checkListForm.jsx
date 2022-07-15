import React, {useState} from "react";
import {v4 as uuidv4} from "uuid";


const CheckListForm = (props) => {
    const [input, setInput] = useState(props.edit ? props.edit.item_text : '');

    const handleChange = e => {
        setInput(e.target.value);
    };

    const handleSubmit = e => {
        e.preventDefault();
        if (props.edit) {
            props.onSubmit({
                item_id: props.edit.item_id,
                item_text: input,
                item_is_complete: props.edit.item_is_complete
            });
        } else {
            props.onSubmit({
                item_id: uuidv4(),
                item_text: input,
                item_is_complete: false
            });
        }
        setInput('');
    };

    return (
        <form className="checklist-form" onSubmit={handleSubmit}>
            { props.edit ? (
                    <React.Fragment>
                        <input
                            className="checklist-input edit"
                            type="text"
                            placeholder="Update item"
                            value={input}
                            name="text"
                            onChange={handleChange}
                            // ref={inputRef}
                            autoFocus
                        />
                        <button className="button is-link is-light  mdi mdi-circle-edit-outline"/>
                    </React.Fragment>
            ) : (
                <React.Fragment>
                    <input
                    className="checklist-input"
                    type="text"
                    placeholder="Add an item... "
                    value={input}
                    name="text"
                    onChange={handleChange}
                    autoFocus
                    />
                    <button className="button is-link is-light  mdi mdi-plus"/>
                </React.Fragment>
                )
            }
        </form>
    );
}

export default CheckListForm;