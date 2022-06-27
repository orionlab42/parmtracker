import React, {useEffect, useRef, useState} from "react";

const CheckListForm = (props) => {
    const [input, setInput] = useState(props.edit ? props.edit.text : '');
    const [id, setId] = useState(0);
    // const inputRef = useRef(null);

    // useEffect(() => {
    //     inputRef.current.focus()
    // });

    const giveId = () => {
        setId(id + 1);
        return id;
    };

    const handleChange = e => {
        setInput(e.target.value);
    };

    const handleSubmit = e => {
        e.preventDefault();
        props.onSubmit({
            id: giveId(),
            text: input,
            isComplete: false
        });
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
                        />
                        <button className="button is-link is-light  mdi mdi-circle-edit-outline"></button>
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
                    />
                    <button className="button is-link is-light  mdi mdi-plus"></button>
                </React.Fragment>
                )
            }
        </form>
    );
}

export default CheckListForm;