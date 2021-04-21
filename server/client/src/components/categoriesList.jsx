import React, {Component} from "react";

class CategoriesList extends Component {
    render() {
        const {categories} = this.props;
        return (
            <div className="field is-grouped is-grouped-multiline category-table">
                { categories.map(category =>
                    <div className="control" key={category.id}>
                        <div className="tags has-addons">
                            <span style={{backgroundColor: category.category_color}} className="tag">
                                <span className={["mdi", category.category_icon].join(" ")}/>
                                {category.category_name}
                            </span>
                            <button
                                className="tag is-delete"
                                /* eslint-disable-next-line no-restricted-globals */
                                onClick={() => confirm("Are you sure you want to delete this category?") ? this.props.onDelete(category.id) : ""}/>
                        </div>
                    </div>)
                }
            </div>
        );
    }
}


export default CategoriesList;