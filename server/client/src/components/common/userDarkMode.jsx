import React, {useEffect, useState} from "react";

const UserDarkMode = (props) => {
    const [darkModeToggle, setDarkModeToggle] = useState(false);

    useEffect(() => {
        setDarkModeToggle(props.user.dark_mode);
    }, [props.user.dark_mode]);


    const handleDarkModeToggle = () => {
        let darkModeToggleChanged;
        darkModeToggleChanged = !darkModeToggle;
        setDarkModeToggle(darkModeToggleChanged);
    };

    return (
        <div>
            <h4 className="title is-5 center-text settings-title">Dark mode:</h4>
            <div className="dark-mode-toggle settings-body">
                <input type="checkbox" id="switch" onChange={props.onChange} onClick={handleDarkModeToggle} checked={darkModeToggle}/>
                <div className="toggle-body">
                    <div className="toggle-container">
                        <label htmlFor="switch">
                            <div className="toggle"/>
                            <div className="names">
                                <p className="light">Light</p>
                                <p className="dark">Dark</p>
                            </div>
                        </label>
                    </div>
                </div>
            </div>
        </div>
    );
};

export default UserDarkMode;