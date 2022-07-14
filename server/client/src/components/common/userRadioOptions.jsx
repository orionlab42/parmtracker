import React, {useEffect, useState} from "react";
import {getUsersOfNote} from "../../services/noteService";
import {getUsers} from "../../services/userService";


const UserRadioOptions = ({user, note, onUserShare}) => {
    const [users, setUsers] = useState([]);
    const [usersToShare, setUsersToShare] = useState([]);

    useEffect(() => {
        getNoteUsersToShare().then();
    }, [users]);

    useEffect(() => {
        async function getALlUsers() {
            const { data: users } = await getUsers();
            if (users != null) {
                setUsers(users);
            }
        }
        getALlUsers();
        getNoteUsersToShare().then();
    }, [user]);

    console.log("users", users)
    console.log("usersToShare", usersToShare)
    const getNoteUsersToShare = async () => {
        const { data: usersOfNote } = await getUsersOfNote(note.note_id);
        if (usersOfNote != null) {
            let usersForShare = [];
            for (let i = 0; i < users.length; i++) {
                let notFound = true
                for (let j = 0; j < usersOfNote.length; j++) {
                    if (users[i].user_id === usersOfNote[j]) {
                        notFound = false;
                        break;
                    }
                }
                if (notFound) {
                    usersForShare.push(users[i]);
                }
            }
            setUsersToShare(usersForShare);
        }
    }

    return (
        <div className="users-checkbox">
            {usersToShare && usersToShare.map(user =>
                <div key={user.user_id}>
                    <input id={user.user_id}
                           className="check-with-label"
                           type="checkbox"
                           name="answer"
                           onClick={() => onUserShare(user.user_id)}/>
                    <label className="checkbox label-for-check" htmlFor={user.user_id}>{user.user_name}</label>
                </div>)}
        {!usersToShare.length && <p>Already shared with everyone.</p>}
        </div>
    );
};

export default UserRadioOptions;