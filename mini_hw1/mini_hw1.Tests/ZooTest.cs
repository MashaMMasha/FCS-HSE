using JetBrains.Annotations;
using Microsoft.VisualStudio.TestTools.UnitTesting;
namespace mini_hw1.Tests;

[TestClass]
[TestSubject(typeof(Zoo))]
public class ZooTest
{
    [TestMethod]
    public void TestZoo()
    {
        var clinic = new VeterinaryClinic();
        var zoo = new Zoo(clinic);
        IAnimal wolf = new Wolf("Wolf", 5, 19, 10, 1, 20);
        IAnimal rabbit = new Rabbit("Rabbit", 5, 3, 10, 2, 10);
        Assert.IsFalse(zoo.AddAnimal(wolf));
        Assert.IsTrue(zoo.AddAnimal(rabbit));
        IThing computer = new Computer(3, "computer");
        Assert.IsTrue(zoo.AddThing(computer));
        zoo.PrintContactZooAnimals();
        zoo.PrintInventory();
        zoo.PrintReport();
    }
}