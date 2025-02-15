namespace mini_hw1;

public class VeterinaryClinic
{
    public bool CheckHealth(IAnimal animal)
    {
        if (animal.IsHealthy() || Cure(animal))
        {
            Console.WriteLine($"Животное {animal.Name} здорово и может быть принято в зоопарк");
        }
        else
        {
            Console.WriteLine($"Животное {animal.Name} не здорово.");
        }
        return animal.IsHealthy();
    }

    bool Cure(IAnimal animal)
    {
        if (animal is IHerbo)
        {
            if (animal is Monkey)
            {
                Console.WriteLine("Вы можете покормить обезьяну бананами и вылечить ее. Хотите это сделать? Если да введите 1.");
            }

            if (animal is Rabbit)
            {
                Console.WriteLine("Вы можете покормить кролика морковкой и вылечить ее. Хотите это сделать? Если да введите 1.");
            }
            var choice = Console.ReadLine();
            if (choice == "1")
            {
                Console.WriteLine($"Животное {animal.Name} теперь здорово и может быть принято в зоопарк");
                animal.GiveFood();
                return true;
            }
        }
        return false;
    }
}