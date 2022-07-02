import React, {useState} from "react";
// import React, {useEffect, useRef, useState} from "react";

const CheckListForm = (props) => {
    const [input, setInput] = useState(props.edit ? props.edit.text : '');
    // const inputRef = useRef(null);

    // useEffect(() => {
    //     inputRef.current.focus()
    // });

    const handleChange = e => {
        setInput(e.target.value);
    };

    const handleSubmit = e => {
        e.preventDefault();
        if (props.edit) {
            props.onSubmit({
                id: props.edit.id,
                text: input,
                isComplete: props.edit.isComplete
            });
        } else {
            props.onSubmit({
                id: props.newIdItem,
                text: input,
                isComplete: false
            });
        }
        setInput('');
        props.increaseIdItem(props.newIdItem + 1);
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
                    // ref={inputRef}
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