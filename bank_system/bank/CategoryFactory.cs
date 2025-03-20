namespace BankSystem;

public class CategoryFactory
{
    private Dictionary<Guid, Category> _categories = new();
    public Guid CreateCategory(string name, OperationType type)
    {
        var category = new Category(name, type);
        _categories.Add(category.id, category);
        return category.id;
    }
    public bool RenameCategory(Guid id, string newName)
    {
        if (!_categories.ContainsKey(id))
        {
            return false;
        }
        _categories[id].name = newName;
        return true;
    }
    public bool DeleteCategory(Guid id)
    {
        if (!_categories.ContainsKey(id))
        {
            return false;
        }
        _categories.Remove(id);
        return true;
    }
    public List<Category> GetAllCategories()
    {
        return new List<Category>(_categories.Values);
    }
}
