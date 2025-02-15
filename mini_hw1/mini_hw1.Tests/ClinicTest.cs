using JetBrains.Annotations;
using Microsoft.VisualStudio.TestTools.UnitTesting;
namespace mini_hw1.Tests;

[TestClass]
[TestSubject(typeof(VeterinaryClinic))]
public class ClinicTest
{
    [TestMethod]
    public void TestCheckHealth()
    {
        VeterinaryClinic clinic = new VeterinaryClinic();
        IAnimal wolf = new Wolf("Wolf", 5, 19, 10, 1, 20);
        Assert.IsFalse(clinic.CheckHealth(wolf));
        IAnimal rabbit = new Rabbit("Rabbit", 5, 3, 10, 1, 10);
        Assert.IsTrue(clinic.CheckHealth(rabbit));
    }
}